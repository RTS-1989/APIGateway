package main

import (
	"github.com/RTS-1989/go-api-gateway/pkg/auth"
	"github.com/RTS-1989/go-api-gateway/pkg/comment"
	"github.com/RTS-1989/go-api-gateway/pkg/config"
	"github.com/RTS-1989/go-api-gateway/pkg/gonews"
	"github.com/RTS-1989/go-api-gateway/pkg/models"
	"github.com/RTS-1989/go-api-gateway/pkg/textParser"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	in := make(chan models.CreateCommentRequestBody)
	out := make(chan models.CreateCommentRequestBody)
	errChan := make(chan error)
	censoredWords := strings.Split(c.Censored, ",")

	cd := textParser.New(in, out, errChan, censoredWords)
	cd.Run()

	authSvc := *auth.RegisterRoutes(r, &c)
	gonews.RegisterRoutes(r, &c, &authSvc)
	comment.RegisterRoutes(r, &c, &authSvc, in, out)

	err = r.Run(c.Port)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
