package grpcclient // import "application/components/grpc-client"

//import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"context"
	"io"
	"sync/atomic"
	"time"

	grpcTypes "git.webdesk.ru/wd/kit/models/grpc/types"

	grpc "google.golang.org/grpc"
	grpcCredentials "google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

// PingSimple Вызов потокового асинхронного пинг метода GRPC
func (cpn *impl) PingStream(address string, insecure bool, count uint64) (err error) {
	var (
		ctx      context.Context
		cfn      context.CancelFunc
		srCtx    context.Context
		srCfn    context.CancelFunc
		opts     []grpc.DialOption
		con      *grpc.ClientConn
		client   grpcTypes.PingStreamClient
		md       grpcMetadata.MD
		stream   grpcTypes.PingStream_PingStreamClient
		n        uint64
		now      time.Time
		beg      time.Time
		totalBeg time.Time
		lat      time.Duration
		req      *grpcTypes.PingStreamInc
		rcvCount *atomic.Value
		creds    grpcCredentials.TransportCredentials
	)

	srCtx, srCfn = context.WithCancel(context.Background())
	opts = append(opts, grpc.WithBlock())
	switch insecure {
	case true:
		opts = append(opts, grpc.WithInsecure())
	case false:
		if cpn.Cfg.Configuration().PingCaCert != "" {
			if creds, err = grpcCredentials.NewClientTLSFromFile(
				cpn.Cfg.Configuration().PingCaCert,
				cpn.Cfg.Configuration().PingHostname); err != nil {
				log.Fatalf("open and reading ca file %q, error %s", cpn.Cfg.Configuration().PingCaCert, err)
				return
			}
		} else {
			creds = grpcCredentials.NewClientTLSFromCert(nil, cpn.Cfg.Configuration().PingHostname)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// Установка соединения
	if con, err = grpc.Dial(address, opts...); err != nil {
		log.Fatalf("can not connect with server %q, error %s", address, err)
	}
	defer func() {
		_ = con.Close()
	}()
	totalBeg = time.Now()
	// Создание клиента GRPC
	client = grpcTypes.NewPingStreamClient(con)
	// Метаданные
	md = grpcMetadata.Pairs("X-Device-Name", "wd-grpc-ping")
	// Контекст с метаданными
	ctx = grpcMetadata.NewOutgoingContext(context.Background(), md)
	// Прерывание контекста
	ctx, cfn = context.WithCancel(ctx)
	defer cfn()
	// Вызов потокового метода
	if stream, err = client.PingStream(ctx); err != nil {
		log.Fatalf("call PingStream() error %s", err)
		return
	}
	defer func() {
		if err = stream.CloseSend(); err != nil {
			log.Errorf("call CloseSend() error %s", err)
		}
	}()
	// Получение данных
	rcvCount = new(atomic.Value)
	rcvCount.Store(uint64(0))
	go cpn.pingStreamReceiver(srCtx, stream, rcvCount)
	defer func() { srCfn() }()
	// Отправка данных
	for {
		if stream.Context().Err() != nil {
			break
		}
		if count != 0 {
			if n >= count {
				break
			}
		}
		n++
		// Подготовка запроса
		now = time.Now()
		req = &grpcTypes.PingStreamInc{
			Id: n,
			CreateAt: &grpcTypes.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		}
		beg = time.Now()
		switch err = stream.Send(req); err {
		case nil:
		case io.EOF:
			log.Noticef("- канал отправки и получения данных закрыт")
			err = nil
			break
		default:
			log.Errorf("- send error: %s", err)
			break
		}
		lat = time.Since(beg)
		log.Noticef("-[%20d] ping ok, асинхронно получено с сервера ping-ов: [%20d], latency: %v", n, rcvCount.Load().(uint64), lat)
	}
	now, lat = time.Now(), time.Since(totalBeg)
	log.Noticef("общее время выполнения: %v [%q - %q]", lat, totalBeg.Format("02.01.2006 15:04:05.999999999"), now.Format("02.01.2006 15:04:05.999999999"))

	return
}

func (cpn *impl) pingStreamReceiver(ctx context.Context, stream grpcTypes.PingStream_PingStreamClient, count *atomic.Value) {
	var (
		err error
		rsp *grpcTypes.PingStreamOut
	)

	defer func() {
		log.Noticef("- горутина получения данных с сервера завершилась")
	}()
	for {
		if ctx.Err() != nil {
			break
		}
		if stream.Context().Err() != nil {
			break
		}
		// Получение сообщения
		if rsp, err = stream.Recv(); err != nil {
			break
		}
		if rsp.Id != 0 {
			count.Store(count.Load().(uint64) + 1)
		}
	}
}
