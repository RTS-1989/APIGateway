package routes

import (
	"context"
	"github.com/RTS-1989/go-api-gateway/pkg/comment/pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Comments(ctx *gin.Context, c pb.CommentServiceClient) {

	newsId, _ := strconv.ParseUint(ctx.Param("news_id"), 10, 64)

	res, err := c.CommentsByNews(context.Background(), &pb.CommentsByNewsRequest{
		NewsId: newsId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
