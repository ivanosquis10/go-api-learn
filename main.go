package main

import (
	"go_api_learn/controllers"
	"go_api_learn/initializers"
	"go_api_learn/middlwares"
	"os"

	"github.com/gin-gonic/gin"
)


func init() { 
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}


func main() {
	router := gin.Default()

	// PING
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong!!",
		})
	})

	// AUTH ROUTES
	auth := router.Group("/auth")
	auth.POST("/signup", controllers.HandleSignUp)
	auth.POST("/login", controllers.HandleSignIn)
	auth.GET("/session", middlwares.RequireAuth, controllers.HandleUserSession)
	auth.POST("/logout", controllers.HandleLogOut)

	// USER ROUTES
	usersController := controllers.UserController{}
	user := router.Group("/users")
	user.GET("/", usersController.HandleGetAllUsers)
	user.GET("/:id", usersController.GetUser)
	// user.GET("/:id", usersController.HandleGetOneUser)
	// user.POST("/", usersController.HandleCreateUser)
	// user.PUT("/:id", usersController.HandleUpdateUser)
	// user.DELETE("/:id", usersController.HandleDeleteUser)

	router.Run(
		":" + os.Getenv("PORT"),
	)
}