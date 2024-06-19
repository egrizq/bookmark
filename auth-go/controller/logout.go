package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	session, _ := store.Get(ctx.Request, "sessions")
	session.Options.MaxAge = -1
	session.Save(ctx.Request, ctx.Writer)

	ctx.JSON(http.StatusOK, "Logout")
}
