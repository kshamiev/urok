package ping // import "application/controllers/grpc/ping"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"application/modules/grpc"
)

// Interface is an interface of package
type Interface grpc.Implementation

// impl is an implementation of package
type impl struct {
}
