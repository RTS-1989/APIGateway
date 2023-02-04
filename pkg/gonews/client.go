package gonews

import (
	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"github.com/RTS-1989/go-api-gateway/pkg/gonews/middleware"
	"github.com/RTS-1989/go-api-gateway/pkg/gonews/pb"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client pb.GoNewsServiceClient
}

func InitServiceClient(c *config.Config) pb.GoNewsServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.GoNewsSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		log.Println("Could not connect:", err)
	}

	return pb.NewGoNewsServiceClient(cc)
}
