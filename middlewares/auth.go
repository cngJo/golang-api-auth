package middlewares

import (
	"github.com/cngJo/golang-api-auth/auth"
	"github.com/cngJo/golang-api-auth/database"
	"github.com/cngJo/golang-api-auth/models"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.GetHeader("Authorization")
		tokenString := authorizationHeader[len("Bearer "):]
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "Authorization header required"})
			context.Abort()
			return
		}
		token, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		var user models.User

		claims := token.Claims.(*auth.JWTClaim)
		record := database.Instance.Where("email = ?", claims.Email).First(&user)
		if record.Error != nil {
			context.JSON(401, gin.H{"error": "User not found"})
			context.Abort()
			return
		}

		context.Set("user", user)

		context.Next()
	}
}
