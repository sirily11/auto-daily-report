package controllers

import (
	"github.com/gin-gonic/gin"
)

// AuthControllerInterface is an interface for AuthController
type AuthControllerInterface interface {
	SignUp(ctx *gin.Context)
}

// AuthController is a struct for AuthController
type AuthController struct {
}

func (c *AuthController) RegisterRoutes(router *gin.RouterGroup) {

}
