package controller

import (
	"auth-go/helpers"
	"auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	// delete session
	session, err := helpers.STORE.Get(ctx.Request, "session")
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	session.Options.MaxAge = -1

	if err := session.Save(ctx.Request, ctx.Writer); err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// response 200
	model.Response(ctx, http.StatusOK, "you're logout")
}
