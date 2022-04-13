package main

import (
	"context"
	"errors"
	"net/http"

	context_os "github.com/pmpavl/lgen-context-os"
	log "github.com/pmpavl/lgen-log"
	"github.com/pmpavl/lgen/app"
)

const ServiceName string = "lgen"

func main() {
	logger := log.For(ServiceName)
	ctx := context_os.Context(context.Background())
	a := app.Get(logger)

	if err := a.Run(ctx); !errors.Is(err, nil) &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, http.ErrServerClosed) {
		logger.Error().Err(err).Msg("error")
	}

	logger.Info().Msg("shutdown service")
}
