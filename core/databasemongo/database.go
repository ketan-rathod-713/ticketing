package databasemongo

import (
	"context"
	"time"

	"github.com/ketan-rathod-713/ticketing/core/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize(config *configs.Config) (*mongo.Client, error) {
	// connect to mongodb database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// instead of localhost connect to auth-mongo service
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DbUrl))

	return client, err
}
