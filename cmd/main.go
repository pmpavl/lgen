package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/pmpavl/lgen/app"
	"github.com/pmpavl/lgen/pkg/contextOS"
	"github.com/pmpavl/lgen/pkg/log"
)

const ServiceName string = "lgen"

func main() {
	logger := log.For(ServiceName)
	ctx := contextOS.Get(context.Background())
	a := app.Get(logger)

	if err := a.Run(ctx); !errors.Is(err, nil) &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, http.ErrServerClosed) {
		logger.Error().Err(err).Msg("error")
	}

	logger.Info().Msg("shutdown service")
}
