package controllers

import (
	"github.com/cngJo/golang-api-auth/database"
	"github.com/cngJo/golang-api-auth/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(400, gin.H{"error": "Error hashing password"})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(400, gin.H{"error": record.Error})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
