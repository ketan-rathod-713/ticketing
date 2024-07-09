package api

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ketan-rathod-713/ticketing/authentication/dto"
	"github.com/ketan-rathod-713/ticketing/core/models"
)

func GetFormattedErrors(err error) []string {
	var errors = make([]string, 0)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println("invalid validation error")
			return errors
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%v:%v", err.Field(), err.Error()))
		}
	}

	return errors
}

func (a *api) Health(ctx *gin.Context) {
	log.Println("health route called")

	ctx.JSON(200, gin.H{
		"message": "authentication service running",
	})
}

func (a *api) Signup(ctx *gin.Context) {
	// get signup data from user
	var signupReq dto.SignupReq
	err := ctx.Bind(&signupReq)

	if err != nil {
		ctx.JSON(400, models.GetResponse("error", nil, "invalid request"))
		ctx.Abort()
		return
	}

	// TODO: add validations using validator package
	err = a.Validator.Struct(&signupReq)

	if err != nil {
		errors := GetFormattedErrors(err)
		ctx.JSON(200, models.GetResponse("error", nil, errors...))
		ctx.Abort()
		return
	}

	// calling authservice to store user data to database
	userDB, err := a.authService.Signup(signupReq.EmailId, signupReq.Password)

	if err != nil {
		ctx.JSON(500, models.GetResponse("error", nil, fmt.Sprintf("internal error, %v", err.Error())))
		ctx.Abort()
		return
	}

	ctx.JSON(200, models.GetResponse("success", dto.SignupRes{Success: true, Id: userDB.Id.Hex()}, "user model fetched"))
}

func (a *api) Signin(ctx *gin.Context) {
	// get signin data
	var signinReq dto.SigninReq
	err := ctx.Bind(&signinReq)
	if err != nil {
		a.Logger.Infof("invalid signin request %v", err.Error())
		ctx.JSON(400, models.GetResponse("error", nil, "invalid request"))
		ctx.Abort()
		return
	}

	// get user info
	user, err := a.authService.GetUser(signinReq.EmailId)
	if err != nil {
		ctx.JSON(500, models.GetResponse("error", nil, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	// match passwords
	if user.Password != signinReq.Password {
		a.Logger.Infof("password doesn't matching userPassword: %v, signinReqPassword: %v", user.Password, signinReq.Password)
		ctx.JSON(401, models.GetResponse("error", nil, "no account matching given request"))
		ctx.Abort()
		return
	}

	// TODO: set cookie in response
	// generate jwt token

	tokenStr, err := a.JwtHelper.GenerateToken(user.EmailId, "user", time.Hour)
	if err != nil {
		a.Logger.Infof("error generating authentication token %v", err)
		ctx.JSON(401, models.GetResponse("error", nil, "error generating authentication token", err.Error()))
		ctx.Abort()
		return
	}

	ctx.SetCookie("token", tokenStr, 10000, "/", "ticketing.dev.ketan", false, false)

	// return userInfo
	ctx.JSON(200, models.GetResponse("success", dto.SigninRes{Success: true}, "signed in"))
}

// get user by emailId for admin
func (a *api) GetUserByEmailId(ctx *gin.Context) {

}

// get current user based on jwt token
func (a *api) GetCurrentUser(ctx *gin.Context) {
	// get token from the client
	token, err := ctx.Cookie("token")
	if err != nil {
		a.Logger.Info("error getting cookie")
		ctx.JSON(200, models.GetResponse("error", nil, "error getting cookie", err.Error()))
		ctx.Abort()
		return
	}

	// verify the token
	userClaims, err := a.JwtHelper.ParseAndValidateToken(token)
	if err != nil {
		a.Logger.Info("jwt token validation failed. please generate new token.")
		a.Logger.Info(err)
		ctx.JSON(200, models.GetResponse("error", nil, "jwt token validation failed. please generate new token.", err.Error()))
		ctx.Abort()
		return
	}

	// get the current user from database
	userInfo, err := a.authService.GetCurrentUser(userClaims.EmailId)
	if err != nil {
		a.Logger.Info("error getting current user %v", err)
		ctx.JSON(200, models.GetResponse("error", nil, "error getting current user", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(200, models.GetResponse("success", userInfo, "successfully fetched userinfo."))
}

// logout // set cookie to empty token = ""
func (a *api) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", 10000, "/", "ticketing.dev.ketan", true, true)

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
