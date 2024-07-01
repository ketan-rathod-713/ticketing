package databasemongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize() (*mongo.Client, error) {
	// connect to mongodb database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// instead of localhost connect to auth-mongo service
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://auth-mongo-srv:27017/authentication"))

	return client, err
}
