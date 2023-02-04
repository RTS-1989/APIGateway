package comment

import (
	"fmt"
	"github.com/RTS-1989/go-api-gateway/pkg/comment/middleware"
	"github.com/RTS-1989/go-api-gateway/pkg/models"

	"github.com/RTS-1989/go-api-gateway/pkg/comment/pb"
	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client     pb.CommentServiceClient
	inComment  chan models.CreateCommentRequestBody
	outComment chan models.CreateCommentRequestBody
}

func InitServiceClient(c *config.Config) pb.CommentServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCommentServiceClient(cc)
}
