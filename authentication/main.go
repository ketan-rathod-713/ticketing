package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ketan-rathod-713/ticketing/authentication/api"
	"github.com/ketan-rathod-713/ticketing/core/configs"
	"github.com/ketan-rathod-713/ticketing/core/databasemongo"
	coreLogger "github.com/ketan-rathod-713/ticketing/core/logger"
)

func main() {
	logger := coreLogger.New()
	logger.Infof("logger initialized")

	// load env variables
	config, err := configs.LoadConfigFronEnvFile()
	if err != nil {
		panic(fmt.Sprintf("unable to load configuration %v", err.Error()))
	}
	logger.Infof("configuration loaded from environment variable %v", config)

	// connect to database
	logger.Info("trying to connect to database")
	client, err := databasemongo.Initialize(config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	logger.Infof("connected to mongodb database")

	logger.Info("initialize routes")
	authApi := api.NewApi(client, logger, config)
	r := authApi.InitializeRoutes()

	logger.Info("authentication service running on port 3000")
	http.ListenAndServe(":3000", r)
}
