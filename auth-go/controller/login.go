package controller

import (
	"auth-go/helpers"
	"auth-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	log.Println("login begin")

	// bind json
	user, err := helpers.BindJSON(ctx)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// check username is exist or not
	userPasswordFromDB, err := helpers.IsUsernameExistForLogin(user.Username)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(userPasswordFromDB), []byte(user.Password))
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	}

	// set session
	if err := helpers.SetSession(ctx, user.Username); err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// return 200
	model.Response(ctx, http.StatusOK, user.Username)
	log.Println("login success")
}
