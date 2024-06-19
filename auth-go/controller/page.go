package controller

import (
	"auth-go/database"
	"auth-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Page(ctx *gin.Context) {
	var allData model.User
	database.DB.First(&allData).Scan(&allData)
	log.Println("page data:", allData)

	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get username from context"})
		return
	}

	response := make(map[string]interface{})
	response["username"] = username
	response["data "] = allData

	ctx.JSON(http.StatusOK, response)
}
