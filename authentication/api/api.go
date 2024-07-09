package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ketan-rathod-713/ticketing/authentication/service"
	"github.com/ketan-rathod-713/ticketing/core/configs"
	"github.com/ketan-rathod-713/ticketing/core/jwthelper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type api struct {
	authService service.Service
	JwtHelper   jwthelper.JWTHelper
	Logger      *zap.SugaredLogger
	Validator   *validator.Validate
	Config      *configs.Config
}

func NewApi(client *mongo.Client, logger *zap.SugaredLogger, config *configs.Config) *api {
	return &api{
		authService: service.New(client.Database("authentication"), logger),
		JwtHelper:   jwthelper.New("secret"),
		Logger:      logger,
		Validator:   validator.New(validator.WithRequiredStructEnabled()),
		Config:      config,
	}
}

func (a *api) InitializeRoutes() *gin.Engine {
	r := gin.Default()

	// users routes
	r.GET("/health", a.Health)

	// user specific routes
	r.POST("/signup", a.Signup)
	r.POST("/signin", a.Signin)
	r.GET("/currentuser", a.GetCurrentUser)
	r.GET("/logout", a.Logout)

	// admin routes
	// TODO: delete user, get all users, get user info by email id

	return r
}

// TODO: authorization middleware
