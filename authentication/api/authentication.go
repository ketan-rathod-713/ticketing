package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ketan-rathod-713/ticketing/authentication/dto"
	"github.com/ketan-rathod-713/ticketing/core/models"
)

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
	// user info validations
	if signupReq.EmailId == "" || signupReq.FirstName == "" || signupReq.LastName == "" || signupReq.Password == "" {
		ctx.JSON(400, models.GetResponse("error", nil, "invalid request"))
		ctx.Abort()
		return
	}

	if len(signupReq.Password) < 5 {
		ctx.JSON(400, models.GetResponse("error", nil, "invalid request, length of password must be more then 5"))
		ctx.Abort()
		return
	}

	// calling authservice to store user data to database
	userDB, err := a.authService.Signup(signupReq.EmailId, signupReq.Password, signupReq.FirstName, signupReq.LastName)

	if err != nil {
		ctx.JSON(500, models.GetResponse("error", nil, "internal error, unable to save user data to database"))
		ctx.Abort()
		return
	}

	ctx.JSON(200, models.GetResponse("success", userDB, "user model fetched"))
}

func (a *api) Signin(ctx *gin.Context) {
	// get signin data
	var signinReq dto.SigninReq
	err := ctx.Bind(&signinReq)
	if err != nil {
		ctx.JSON(400, models.GetResponse("error", nil, "invalid request"))
		ctx.Abort()
		return
	}

	// get user info
	user, err := a.authService.GetUser(signinReq.EmailId)
	if err != nil {
		ctx.JSON(500, models.GetResponse("error", nil, "internal server error"))
		ctx.Abort()
		return
	}

	// match passwords
	if user.Password != signinReq.Password {
		ctx.JSON(401, models.GetResponse("error", nil, "no account matching given request"))
		ctx.Abort()
		return
	}

	// TODO: set cookie in response
	// generate jwt token
	

	// return userInfo
	ctx.JSON(200, models.GetResponse("success", dto.SigninRes{Success: true}, "signed in"))
}

// get user by emailId for admin
func (a *api) GetUserByEmailId(ctx *gin.Context) {

}

// get current user based on jwt token
func(a *api) GetCurrentUser(ctx *gin.Context){

}

// logout // set cookie to empty token = ""
func (a *api) Logout(ctx *gin.Context) {

}
