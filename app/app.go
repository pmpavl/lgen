package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pmpavl/lgen/handler"
	"github.com/pmpavl/lgen/pkg/storage"
	"github.com/pmpavl/lgen/pkg/tex"
	"github.com/pmpavl/lgen/resource"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type App struct {
	logger *zerolog.Logger
}

func Get(l *zerolog.Logger) *App {
	return &App{logger: l}
}

func (a *App) Run(ctx context.Context) error { // TODO: Нужно проинициализировать public/gen и public/tex
	res := resource.Get(ctx)
	storage := storage.Get(ctx, res.Mongo)
	tex := tex.Get()
	h := handler.Get(storage, tex)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", res.Env.ServiceHTTPPort),
		Handler: a.ginRouter(h),
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.logger.Info().Int("port", res.Env.ServiceHTTPPort).Msg("start http")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Err(err).Msg("listen and serve")

			return err
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		return srv.Shutdown(ctx)
	})

	return group.Wait()
}
