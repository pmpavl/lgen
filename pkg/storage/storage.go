package storage

import (
	"context"

	"github.com/pmpavl/lgen/core/constant"
	"github.com/pmpavl/lgen/pkg/log"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

const PackageName string = "storage"

type Storage struct {
	logger       *zerolog.Logger
	collTask     *mongo.Collection
	collTemplate *mongo.Collection
}

func Get(ctx context.Context, client *mongo.Client) *Storage {
	mongo := client.Database(constant.MongoDatabaseName)

	return &Storage{
		logger:       log.For(PackageName),
		collTask:     mongo.Collection(constant.MongoCollectionTaskName),
		collTemplate: mongo.Collection(constant.MongoCollectionTemplateName),
	}
}
