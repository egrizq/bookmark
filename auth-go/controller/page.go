package controller

import (
	"auth-go/helpers"
	"auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Page(ctx *gin.Context) {
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil || username == "" {
		model.Response(ctx, http.StatusUnauthorized, "unauthorized user")
		return
	}

	model.Response(ctx, http.StatusOK, username)
}
