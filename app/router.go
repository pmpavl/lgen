package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pmpavl/lgen/app/middleware"
	"github.com/pmpavl/lgen/pkg/log"
)

type Handler interface {
	GeneratePing(c *gin.Context)
}

func (a *App) ginRouter(h Handler) *gin.Engine {
	router := gin.New()
	router.Use(
		middleware.CORS(),
		middleware.Logger(log.For("http-server")),
	)
	router.Static("/static", "./public")

	router.POST("/app/v1/generate/ping", h.GeneratePing)

	return router
}
