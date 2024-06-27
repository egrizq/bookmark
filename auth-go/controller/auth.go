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

func Login(ctx *gin.Context) {
	// bind json
	user, err := helpers.BindJSON(ctx)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, "invalid request")
		return
	}

	// check username is exist or not
	userPasswordFromDB, err := helpers.IsUsernameExistForLogin(user.Username)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "username atau password salah")
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(userPasswordFromDB), []byte(user.Password))
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "username atau password salah")
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
	saveUser := model.FormatUser(user)
	if result := database.DB.Create(saveUser); result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError,
			"terjadi kesalahan secara internal saat pembuatan akun")
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
