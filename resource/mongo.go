package resource

import (
	"context"
	"fmt"

	"github.com/pmpavl/lgen/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Resource) getMongo(ctx context.Context) error {
	clientOptions := options.Client().
		SetAppName(r.Env.ServiceName).
		ApplyURI(r.Env.MongoHost)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("mongo connect: %s", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("client ping: %s", err)
	}

	r.Mongo = client

	log.Logger.Info().Msg("get mongo success")

	return nil
}
