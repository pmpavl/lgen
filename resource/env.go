package resource

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	ServiceName     string `env:"SERVICE_NAME, default=lgen"`
	ServiceHTTPPort string `env:"SERVICE_HTTP_PORT, default=:8080"`
	LogLevel        string `env:"LOG_LEVEL, default=info"`
	LogFormat       string `env:"LOG_FORMAT, default=console"`
	GinMode         string `env:"GIN_MODE, default=release"`
	MongoHost       string `env:"MONGO_HOST, required"`
}

func (r *Resource) getEnv(ctx context.Context) error {
	var env Env

	if err := envconfig.Process(ctx, &env); err != nil {
		return fmt.Errorf("envconfig process: %s", err)
	}

	r.Env = &env

	return nil
}
