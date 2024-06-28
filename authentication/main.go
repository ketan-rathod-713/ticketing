package main

import (
	"log"
	"net/http"

	"github.com/ketan-rathod-713/ticketing/authentication/api"
	"github.com/ketan-rathod-713/ticketing/core/databasemongo"
)

func main() {
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
