package resource

import (
	"context"
	"fmt"

	storage "github.com/pmpavl/lgen-storage"
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

	storage, err := storage.Get(ctx, client)
	if err != nil {
		return fmt.Errorf("get storage: %s", err)
	}

	r.Storage = storage

	return nil
}
