package middlewares

import (
	"github.com/cngJo/golang-api-auth/auth"
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
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
