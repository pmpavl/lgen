package contextOS

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pmpavl/lgen/pkg/log"
)

const PackageName string = "context_os"

func Get(ctx context.Context) context.Context {
	logger := log.For(PackageName)
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}

	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)

	go func() {
		select {
		case <-ctx.Done():
		case sig := <-sigs:
			logger.Info().Msgf("got signal %s", sig)
			cancel()
		}
	}()

	return ctx
}
