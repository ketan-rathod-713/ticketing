package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ketan-rathod-713/ticketing/authentication/api"
	"github.com/ketan-rathod-713/ticketing/core/configs"
	"github.com/ketan-rathod-713/ticketing/core/databasemongo"
)

func main() {
	// load env variables
	config, err := configs.LoadConfigFronEnvFile()
	if err != nil {
		panic(fmt.Sprintf("unable to load configuration %v", err.Error()))
	}

	log.Println("Configuration loaded", config)

	// connect to database
	client, err := databasemongo.Initialize()
	if err != nil {
		panic(err)
	}

	authApi := api.NewApi(client)

	r := authApi.InitializeRoutes()

	log.Println("authentication service running on port 3000")
	http.ListenAndServe(":3000", r)
}

// TODO: get env variables from .env file, what we will do in docker container ??
