package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/pmpavl/lgen/app"
	"github.com/pmpavl/lgen/pkg/context_os"
	"github.com/pmpavl/lgen/pkg/log"
)

const ServiceName string = "lgen"

func main() {
	logger := log.For(ServiceName)
	ctx := context_os.Get(context.Background())
	a := app.Get(logger)

	if err := a.Run(ctx); !errors.Is(err, nil) &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, http.ErrServerClosed) {
		logger.Error().Err(err).Msg("error")
	}

	logger.Info().Msg("shutdown service")
}
