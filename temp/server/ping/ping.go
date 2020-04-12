package ping // import "application/controllers/grpc/ping"

//import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"context"
	"fmt"
	"io"
	"math"
	runtimeDebug "runtime/debug"
	"sync"
	"time"

	grpcTypes "git.webdesk.ru/wd/kit/models/grpc/types"
	grpc "google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
	grpcPeer "google.golang.org/grpc/peer"
)

// New creates a new object and return interface
func New() Interface {
	var ping = new(impl)
	return ping
}

// Register Регистрация обработчиков запросов GRPC сервера
func (ping *impl) Register(srv *grpc.Server) {
	grpcTypes.RegisterPingServer(srv, ping)
	grpcTypes.RegisterPingStreamServer(srv, ping)
}

// Ping Проверка доступности GRPC сервера, реализация интерфейса сервера
func (ping *impl) Ping(ctx context.Context, rq *grpcTypes.PingRequest) (ret *grpcTypes.PingResponse, err error) {
	const timeFormat = `02.01.2006 15:04:05.999999999Z07:00`
	var (
		nowAt time.Time
		reqAt time.Time
	)

	defer func() {
		// Восстановление после паники
		if e := recover(); e != nil {
			err = fmt.Errorf("panic recovery:\n%v\n%s", e.(error), string(runtimeDebug.Stack()))
		}
	}()
	// DEBUG
	//log.Debugf("- Ping(), подключение клиента")
	// DEBUG
	nowAt = time.Now()
	if rq.CreateAt != nil && rq.CreateAt.Seconds > 0 && rq.CreateAt.Nanos > 0 {
		reqAt = time.Unix(rq.CreateAt.Seconds, int64(rq.CreateAt.Nanos))
	}
	ret = &grpcTypes.PingResponse{
		Message: fmt.Sprintf("RE: %s (%s)", rq.Message, reqAt.Format(timeFormat)),
		CreateAt: &grpcTypes.Timestamp{
			Seconds: nowAt.Unix(),
			Nanos:   int32(nowAt.Nanosecond()),
		},
	}
	// DEBUG
	//log.Debugf("- Ping(), отключение клиента")
	// DEBUG

	return
}

// PingStream Проверка доступности GRPC сервера в потоковом режиме, реализация интерфейса сервера
func (ping *impl) PingStream(stream grpcTypes.PingStream_PingStreamServer) (err error) {
	var (
		pr     *grpcPeer.Peer
		md     grpcMetadata.MD
		isDone bool
		ok     bool
		req    *grpcTypes.PingStreamInc
		inp    chan *grpcTypes.PingStreamInc
		swg    *sync.WaitGroup
		ctx    context.Context
		cfn    context.CancelFunc
	)

	// Контекст завершения горутины
	ctx, cfn = context.WithCancel(context.Background())
	defer func() {
		// Восстановление после паники
		if e := recover(); e != nil {
			err = fmt.Errorf("panic recovery:\n%v\n%s", e.(error), string(runtimeDebug.Stack()))
			// Остановка горутины
			if cfn != nil {
				cfn()
			}
		}
	}()
	// DEBUG
	//log.Debugf("- PingStream(), подключение потокового клиента")
	// DEBUG
	// Загрузка информации о GRPC клиенте (ip, port, info)
	if pr, ok = grpcPeer.FromContext(stream.Context()); ok {
		// DEBUG
		//log.Debug(debug.DumperString(pr))
		// DEBUG
		_ = pr
	}
	// Загрузка метаданных (заголовки запроса)
	if md, ok = grpcMetadata.FromIncomingContext(stream.Context()); ok {
		// DEBUG
		//log.Debug(debug.DumperString(md))
		// DEBUG
		_ = md
	}
	// Канал потока входящих сообщений
	inp = make(chan *grpcTypes.PingStreamInc)
	defer close(inp)
	// Ожидание завершение горутины через sync.WaitGroup
	swg = new(sync.WaitGroup)
	swg.Add(1)
	// Запуск горутины
	go ping.PingStreamReaderWriter(ctx, stream, inp, swg)
	for {
		if isDone {
			break
		}
		if stream.Context().Err() != nil {
			isDone = true
			continue
		}
		switch req, err = stream.Recv(); err {
		case nil:
		case io.EOF:
			isDone, err = true, nil
			continue
		default:
			err = fmt.Errorf("ping stream receive error: %s", err)
			break
		}
		inp <- req
	}
	// DEBUG
	//log.Debugf("- PingStream(), отключение потокового клиента, ожидание завершения горутины")
	// DEBUG
	cfn()
	swg.Wait()

	return
}

// PingStreamReaderWriter Горутина обслуживания потоковых запросов одного клиента
func (ping *impl) PingStreamReaderWriter(
	ctx context.Context,
	stream grpcTypes.PingStream_PingStreamServer,
	inp <-chan *grpcTypes.PingStreamInc,
	swg *sync.WaitGroup,
) {
	const sendTimeout = time.Second * 2
	var (
		err    error
		req    *grpcTypes.PingStreamInc
		tic    *time.Ticker
		now    time.Time
		count  uint64
		isDone bool
	)

	defer swg.Done()
	tic = time.NewTicker(sendTimeout)
	defer tic.Stop()
	// DEBUG
	//log.Debugf("- PingStream(), запуск горутины обслуживания потокового клиента")
	//defer log.Debugf("- PingStream(), завершение горутины обслуживания потокового клиента")
	// DEBUG
	for {
		if isDone {
			break
		}
		if count < math.MaxUint64/2 {
			count = math.MaxUint64
		}
		select {
		// Прерывание обмена данными
		case <-ctx.Done():
			isDone = true
			continue
		// Канал входящих сообщений
		case req = <-inp:
			if req == nil {
				isDone = true
				continue
			}
			// DEBUG
			//log.Debug(debug.DumperString(req))
			// DEBUG
			if err = stream.Send(&grpcTypes.PingStreamOut{
				Id: req.Id,
				CreateAt: &grpcTypes.Timestamp{
					Seconds: req.CreateAt.Seconds,
					Nanos:   req.CreateAt.Nanos,
				},
			}); err != nil {
				log.Criticalf("ping stream reader writer send message error: %s", err)
				continue
			}
		// Источник событий генерации исходящих сообщений
		case <-tic.C:
			count--
			now = time.Now()
			if err = stream.Send(&grpcTypes.PingStreamOut{
				Id: count,
				CreateAt: &grpcTypes.Timestamp{
					Seconds: now.Unix(),
					Nanos:   int32(now.Nanosecond()),
				},
			}); err != nil {
				log.Criticalf("ping stream reader writer send message error: %s", err)
				continue
			}
		}
	}
}
