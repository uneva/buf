package v1

import (
	"context"
	"net/http"

	transport "github.com/uneva/buf/playground/transport/gin"
)

type GreeterServiceGinServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

func RegisterGreeterServiceGinServer(s *transport.Server, srv GreeterServiceGinServer) {
	r := s.Router()
	r.GET("/v1/hello/:name", _GreeterService_SayHello_Gin_Handler(srv))
}

func _GreeterService_SayHello_Gin_Handler(srv GreeterServiceGinServer) func(ctx transport.Context) error {
	return func(ctx transport.Context) error {
		var in SayHelloRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		out, err := srv.SayHello(ctx, &in)
		if err != nil {
			return err
		}

		return ctx.Reply(http.StatusOK, out)
	}
}
