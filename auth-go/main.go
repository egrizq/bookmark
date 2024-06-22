package main

import (
	"auth-go/controller"
	"auth-go/database"
	"auth-go/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connection()
	helpers.InitSession()

	router := gin.Default()
	router.Use(helpers.CORSMiddleware())

	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)

	app := router.Group("/bookmark")
	app.Use(helpers.CheckSession())

	// testing
	app.GET("/page", controller.Page)

	// bookmarks
	app.POST("/add", controller.NewBookmark)
	app.GET("/get/:category", controller.GetBookmarkByCategory)

	// category
	app.POST("/category/add", controller.NewCategory)

	router.Run(":8000")
}
