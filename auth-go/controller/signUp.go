package controller

import (
	"auth-go/database"
	"auth-go/helpers"
	"auth-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var user model.User

	// bind json
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError,
			"An error occurred while creating your account.")
		return
	}
	log.Println("json data:", user)

	// check username if exist or not
	var checkUser model.User
	result := database.DB.Where("username = ?", user.Username).First(&checkUser)
	if result.RowsAffected == 1 {
		model.Response(ctx, http.StatusBadRequest,
			"Username already exists, please use a different Username")
		return
	}

	// hash password
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPass)

	// insert to database
	if result := database.DB.Create(user); result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError, "An error occurred while creating your account")
		return
	}

	// set jwt
	token, err := helpers.JWTGenerate(user.Username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
	})

	// return 201
	model.Response(ctx, http.StatusCreated, "success")
}
