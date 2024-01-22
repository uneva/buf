package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/uneva/buf/playground/api/greeter/v1"
)

type service struct{}

func (s *service) SayHello(ctx context.Context, req *v1.SayHelloRequest) (*v1.SayHelloResponse, error) {
	return &v1.SayHelloResponse{
		Text: fmt.Sprintf("Hi %s~", req.Name),
	}, nil
}

func main() {
	e := gin.Default()
	e.GET("/openapi", func(ctx *gin.Context) {
		ctx.File("./api/openapi.yaml")
	})
	// srv := transport.NewServer(transport.Engine(e))

	// srv.Router().GET("/openapi", func(ctx transport.Context) error {
	// 	ctx.File("./api/openapi.yaml")
	// 	return nil
	// })

	// v1.RegisterGreeterServiceGinServer(srv, &service{})

	e.Run()
}
