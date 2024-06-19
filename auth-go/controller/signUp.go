package controller

import (
	"auth-go/database"
	"auth-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var registerData model.User

	// bind json
	err := ctx.ShouldBindJSON(&registerData)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError,
			"An error occurred while creating your account.")
		return
	}
	log.Println("json data:", registerData)

	// check username if exist or not
	var checkData model.User
	result := database.DB.Where("username = ?", registerData.Username).First(&checkData)
	if result.RowsAffected == 1 {
		model.Response(ctx, http.StatusBadRequest,
			"Username already exists, please use a different Username")
		return
	}

	// hash password
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), 14)
	registerData.Password = string(hashPass)

	// insert to database
	if result := database.DB.Create(registerData); result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError, "An error occurred while creating your account")
		return
	}

	// return 201
	model.Response(ctx, http.StatusCreated, "success")
}
