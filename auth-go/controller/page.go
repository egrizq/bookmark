package controller

import (
	"auth-go/database"
	"auth-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Page(ctx *gin.Context) {
	session, _ := store.Get(ctx.Request, "sessions")
	if auth, ok := session.Values["validate"].(bool); !ok || !auth {
		ctx.JSON(http.StatusUnauthorized, "there's no sessions")
		return
	}

	var allData model.User
	database.DB.First(&allData).Scan(&allData)
	log.Println("page data:", allData)

	response := make(map[string]interface{})
	response["sessions"] = session.Values["account"]
	response["data "] = allData

	ctx.JSON(http.StatusOK, response)
}
