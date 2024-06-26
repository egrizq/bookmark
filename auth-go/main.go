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
	router.Use(helpers.CORS())

	// public
	router.POST("/login", controller.Login)
	router.POST("/signup", controller.SignUp)

	// public-bookmark
	router.GET("/get/:username/:category", controller.GetBookmarkByCategory)

	// app
	app := router.Group("/bookmark")
	app.Use(helpers.CheckSession())

	// app-page
	app.GET("/page", controller.Page)

	// app-bookmarks
	app.POST("/insert", controller.NewBookmark)
	app.GET("/list", controller.GetListOfCategoryAndNumberOfBookmarks)

	// app-category
	app.POST("/category/insert", controller.InsertNewCategory)
	app.GET("/category/list", controller.GetListCategory)

	// logout
	app.POST("/logout", controller.Logout)

	router.Run(":8000")
}
