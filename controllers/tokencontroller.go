package controllers

import (
	"github.com/cngJo/golang-api-auth/auth"
	"github.com/cngJo/golang-api-auth/database"
	"github.com/cngJo/golang-api-auth/models"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(401, gin.H{"error": "User not found"})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(401, gin.H{"error": "Invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(500, gin.H{"error": "Error generating token"})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{"token": tokenString})
}
