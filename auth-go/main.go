package main

import (
	"auth-go/controller"
	"auth-go/database"
	"auth-go/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connection()

	router := gin.Default()
	router.Use(helpers.CORSMiddleware())

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)

	app := router.Group("/page")
	app.Use(helpers.JWTMiddleware())

	app.GET("/main", controller.Page)

	router.Run(":8000")
}
