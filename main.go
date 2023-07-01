package main

import (
	"github.com/cngJo/golang-api-auth/controllers"
	"github.com/cngJo/golang-api-auth/database"
	"github.com/cngJo/golang-api-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root:root@tcp(127.0.0.1:3306)/app?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.GenerateToken)
			auth.POST("/register", controllers.RegisterUser)
		}
		
		secured := api.Group("").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
