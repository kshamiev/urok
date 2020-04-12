package grpcclient // import "application/components/grpc-client"

import "gopkg.in/webnice/debug.v1"
import "gopkg.in/webnice/log.v2"
import (
	"context"
	"fmt"
	"strings"
	"time"

	grpcTypes "git.webdesk.ru/wd/kit/models/grpc/types"

	grpc "google.golang.org/grpc"
	grpcCredentials "google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

// PingSimple Вызов одиночного синхронного метода GRPC на каждый пинг
func (cpn *impl) PingSimple(address string, insecure bool, count uint64) (err error) {
	const maxRetry = 3
	var (
		n     uint64
		retry uint64
	)

	for {
		if count != 0 {
			if n >= count {
				break
			}
		}
		n++
		for retry = 0; retry < maxRetry; retry++ {
			if err = cpn.pingSimple(address, insecure, n); err == nil {
				break
			}
		}
		if err != nil {
			return
		}
	}

	return
}

func (cpn *impl) pingSimple(address string, insecure bool, n uint64) (err error) {
	const uk = `e442e91a4e198a61bfc8d4ab6e2be9d86690ae118d0bae999336529282260362d748e6ec7c9940642d0956338daa70ca896208a73302a734c68caa4808a6a613`
	var (
		ctx    context.Context
		cfn    context.CancelFunc
		opts   []grpc.DialOption
		con    *grpc.ClientConn
		client grpcTypes.PingClient
		md     grpcMetadata.MD
		now    time.Time
		beg    time.Time
		lat    time.Duration
		req    *grpcTypes.PingRequest
		rsp    *grpcTypes.PingResponse
		creds  grpcCredentials.TransportCredentials
	)

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
	// Создание клиента GRPC
	client = grpcTypes.NewPingClient(con)
	// Метаданные
	md = grpcMetadata.Pairs("X-Device-Name", "wd-grpc-ping")
	// Контекст с метаданными
	ctx = grpcMetadata.NewOutgoingContext(context.Background(), md)
	// Прерывание контекста
	ctx, cfn = context.WithCancel(ctx)
	defer cfn()
	// Подготовка запроса
	now = time.Now()
	req = &grpcTypes.PingRequest{
		Message: uk,
		CreateAt: &grpcTypes.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}
	// Выполнение запроса
	beg = time.Now()
	if rsp, err = client.Ping(ctx, req); err != nil {
		err = fmt.Errorf("- выполнение ping запроса завершилось ошибкой: %s", err)
		return
	}
	lat = time.Since(beg)
	if !strings.Contains(rsp.Message, uk) {
		log.Errorf("- запрос отправлен, ответ получен, но ответ не корректный")
		log.Debug(debug.DumperString(rsp))
		return
	}
	log.Noticef("-[%20d] ping ok, latency: %v", n, lat)

	return
}
