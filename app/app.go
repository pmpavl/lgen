package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/pmpavl/lgen-log"
	"github.com/pmpavl/lgen/app/middleware"
	"github.com/pmpavl/lgen/handler"
	"github.com/pmpavl/lgen/resource"
	"github.com/pmpavl/lgen/tex"
	"github.com/rs/zerolog"
)

type App struct {
	logger *zerolog.Logger
}

func Get(l *zerolog.Logger) *App {
	return &App{logger: l}
}

func (a *App) Run(ctx context.Context) error {
	res := resource.Get(ctx)
	tex := tex.Get()
	h := handler.Get(res.Storage, tex)

	router := gin.New()

	router.Use(middleware.Logger(log.For("http-server")))

	// Mount Handler
	router.POST("/app/generate/leaflet", h.GenerateLeaflet)

	srv := &http.Server{
		Addr:    res.Env.ServiceHTTPPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal().Msgf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	return srv.Shutdown(ctx)
}
