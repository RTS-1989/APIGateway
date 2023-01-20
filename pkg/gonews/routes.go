package gonews

import (
	"github.com/RTS-1989/go-api-gateway/pkg/auth"
	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"github.com/RTS-1989/go-api-gateway/pkg/gonews/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/gonews")
	routesGroup.Use(a.AuthRequired)
	routesGroup.GET("/:n", svc.Posts)
	fullNews := routesGroup.Group("/news_full")
	fullNews.GET("/:news_id", svc.NewsFullDetailed)
	shortNews := routesGroup.Group("/news_short")
	shortNews.GET("/:news_id", svc.NewsShortDetailed)
	filteredNews := routesGroup.Group("/filtered_news")
	filteredNews.GET("/:filter_value", svc.FilterNews)
}

func (svc *ServiceClient) Posts(ctx *gin.Context) {
	routes.Posts(ctx, svc.Client)
}

func (svc *ServiceClient) NewsFullDetailed(ctx *gin.Context) {
	routes.NewsFullDetailed(ctx, svc.Client)
}

func (svc *ServiceClient) NewsShortDetailed(ctx *gin.Context) {
	routes.NewsShortDetailed(ctx, svc.Client)
}

func (svc *ServiceClient) FilterNews(ctx *gin.Context) {
	routes.FilterNews(ctx, svc.Client)
}
