package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ketan-rathod-713/ticketing/authentication/service"
	"github.com/ketan-rathod-713/ticketing/core/jwthelper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type api struct {
	authService service.Service
	JwtHelper   jwthelper.JWTHelper
	Logger      *zap.SugaredLogger
}

func NewApi(client *mongo.Client, logger *zap.SugaredLogger) *api {
	return &api{
		authService: service.New(client.Database("authentication"), logger),
		JwtHelper:   jwthelper.New("secret"),
		Logger:      logger,
	}
}

func (a *api) InitializeRoutes() *gin.Engine {
	r := gin.Default()

	// users routes
	r.GET("/health", a.Health)
	r.POST("/signup", a.Signup)
	r.POST("/signin", a.Signin)
	r.GET("/currentuser", a.GetCurrentUser)
	r.GET("/logout", a.Logout)

	// admin routes
	// TODO: delete user, get all users, get user info by email id

	return r
}
