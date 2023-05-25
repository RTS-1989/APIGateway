package main

import (
	"github.com/RTS-1989/go-api-gateway/pkg/auth"
	"github.com/RTS-1989/go-api-gateway/pkg/censor"
	"github.com/RTS-1989/go-api-gateway/pkg/comment"
	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"github.com/RTS-1989/go-api-gateway/pkg/gonews"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		c.FrontendSvcUrl,
	}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	r.Use(cors.New(corsConfig))

	authSvc := *auth.RegisterRoutes(r, &c)
	gonews.RegisterRoutes(r, &c, &authSvc)
	comment.RegisterRoutes(r, &c, &authSvc)
	censor.RegisterRoutes(r, &c, &authSvc)

	err = r.Run(":" + c.Port)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
