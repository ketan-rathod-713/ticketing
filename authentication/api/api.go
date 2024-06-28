package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ketan-rathod-713/ticketing/authentication/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	authService service.Service
}

func NewApi(client *mongo.Client) *api {
	return &api{
		authService: service.New(client.Database("authentication")),
	}
}

func (a *api) InitializeRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health", a.Health)
	r.GET("/signup", a.Signup)

	return r
}
