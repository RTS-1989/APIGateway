package gonews

import (
	"fmt"

	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"github.com/RTS-1989/go-api-gateway/pkg/gonews/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.GoNewsServiceClient
}

func InitServiceClient(c *config.Config) pb.GoNewsServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.GoNewsSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewGoNewsServiceClient(cc)
}
