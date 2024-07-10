package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ketan-rathod-713/ticketing/authentication/service"
	"github.com/ketan-rathod-713/ticketing/core/configs"
	"github.com/ketan-rathod-713/ticketing/core/jwthelper"
	"github.com/ketan-rathod-713/ticketing/core/models"
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

	rg := r.Group("/authentication")

	// users routes
	rg.GET("/health", a.Health)

	// user specific routes
	rg.POST("/signup", a.Signup)
	rg.POST("/signin", a.Signin)

	// TODO: add token parser middleware to it
	rg.GET("/currentuser", a.GetCurrentUser)
	rg.GET("/logout", a.Logout)

	// Todo: add token parser middleware to it
	rg.GET("/verifyEmail", a.TokenParser, a.VerifyEmail)

	// admin routes
	// TODO: delete user, get all users, get user info by email id

	return r
}

// TODO: authorization middleware
// it will take token and parse it and pass the claims inside the gin context
// middleware token parser
func (a *api) TokenParser(ctx *gin.Context) {
	// take token from cookie
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(200, models.GetResponse("error", nil, "error getting cookie", err.Error()))
		ctx.Abort()
		return
	}

	// parse and validate the token
	userClaims, err := a.JwtHelper.ParseAndValidateToken(token)
	if err != nil {
		ctx.JSON(200, models.GetResponse("error", nil, "jwt token validation failed. please generate new token.", err.Error()))
		ctx.Abort()
		return
	}

	// pass the claims to the context
	ctx.Set("user", userClaims)

	// call next function
	ctx.Next()
}
