package grpc // import "application/controllers/grpc"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"application/controllers/grpc/msg"
	"application/controllers/grpc/ping"
	"application/modules/grpc"
)

var (
	// EchoController grpc controller interface
	EchoController = ping.New()

	// MsgController grpc controller interface
	MsgController = msg.New()
)

// Регистрация всех необходимых GRPC контроллеров
// Если зарегистрировал хотябы один контроллер, роутинг изменяется автоматически и один и тот же сервер
// начинает обрабатывать как REST запросы так и GRPC запросы
func init() {
	grpc.RegistrationGrpcController(EchoController)
	grpc.RegistrationGrpcController(MsgController)
}
