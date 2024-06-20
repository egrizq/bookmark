package controller

import (
	"auth-go/helpers"
	"auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Page(ctx *gin.Context) {
	session, _ := helpers.STORE.Get(ctx.Request, "session")
	username, _ := session.Values["username"].(string)

	model.Response(ctx, http.StatusOK, username)
}
