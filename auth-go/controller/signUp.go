package controller

import (
	"auth-go/database"
	"auth-go/helpers"
	"auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	// bind json
	user, err := helpers.BindJSON(ctx)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// check username if exist or not
	if err = helpers.IsUsernameExistForSignUp(user.Username); err != nil {
		model.Response(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// hash password
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPass)

	// insert to database
	if result := database.DB.Create(user); result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError, "an error occurred while creating your account")
		return
	}

	// set session
	if err := helpers.SetSession(ctx, user.Username); err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// return 201
	model.Response(ctx, http.StatusCreated, "success")
}
