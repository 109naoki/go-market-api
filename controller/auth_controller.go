package controllers

import (
	"gin-market/dto"
	"gin-market/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface{
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AuthController struct{
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController{
	return &AuthController{service: service}
}

func (c *AuthController) SignUp(ctx *gin.Context){
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.SignUp(input.Email, input.Password)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func(c *AuthController)Login(ctx *gin.Context){
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(input.Email, input.Password)
	if err != nil{
		if err.Error() == "User not found"{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}