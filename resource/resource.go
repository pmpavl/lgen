package resource

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/lgen/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
)

type Resource struct {
	Env   *Env
	Mongo *mongo.Client
}

func Get(ctx context.Context) *Resource {
	r := &Resource{}

	if err := r.loadDotEnv(); err != nil {
		log.Logger.Fatal().Err(err).Msg("load dotenv")
	}

	if err := r.getEnv(ctx); err != nil {
		log.Logger.Fatal().Err(err).Msg("get env")
	}

	setLogger(r)
	setGinMode(r)

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := r.getMongo(ctx); err != nil {
			return fmt.Errorf("get mongo: %s", err)
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		log.Logger.Fatal().Err(err).Msg("get resource")
	}

	log.Logger.Info().Msg("get resource success")

	return r
}

func setLogger(r *Resource) {
	logLevel, err := log.ParseLogLevel(r.Env.LogLevel)
	if err != nil {
		log.Logger.Warn().Err(err).Msg("parse log level")
	}

	if err == nil && logLevel != log.LogLevelDefault {
		log.SetGlobalLevel(logLevel)
	}

	logFormat, err := log.ParseLogFormat(r.Env.LogFormat)
	if err != nil {
		log.Logger.Warn().Err(err).Msg("parse log format")
	}

	if err == nil && logFormat != log.LogFormatDefault {
		log.SetGlobalFormat(logFormat)
	}

	return
}

func setGinMode(r *Resource) {
	switch r.Env.GinMode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default: // missmatch gin mode
		log.Logger.Warn().
			Err(fmt.Errorf("no such gin mode: %s", r.Env.GinMode)).
			Msg("set gin mode")

		return
	}

	log.Logger.Info().Msgf("gin mode set to %s", gin.Mode())

	return
}
