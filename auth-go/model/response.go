package model

import "github.com/gin-gonic/gin"

type ResponseType struct {
	StatusCode int
	Message    string
}

func Response(ctx *gin.Context, status int, message string) {
	res := ResponseType{
		StatusCode: status,
		Message:    message,
	}

	ctx.JSON(res.StatusCode, res)
}
