package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (a *api) Health(ctx *gin.Context) {
	log.Println("health route called")

	ctx.JSON(200, gin.H{
		"message": "authentication service running",
	})
}

func (a *api) Signup(ctx *gin.Context) {
	log.Println("signup route called")

	ctx.JSON(200, gin.H{
		"message": "signup succes",
	})
}
func (a *api) Signin(ctx *gin.Context) {

}
func (a *api) GetUser(ctx *gin.Context) {

}
func (a *api) Logout(ctx *gin.Context) {

}
