package censor

import (
	"fmt"
	"github.com/RTS-1989/go-api-gateway/pkg/censor/middleware"
	"github.com/RTS-1989/go-api-gateway/pkg/censor/pb"
	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CensorServiceClient
}

func InitServiceClient(c *config.Config) pb.CensorServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CensoredSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCensorServiceClient(cc)
}
