package controller

import (
	"auth-go/database"
	"auth-go/helpers"
	"auth-go/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_KEY")))
)

func Login(ctx *gin.Context) {

	// bind json
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, "Invalid login credentials.")
		return
	}
	log.Println("data json:", user)

	// check username is exist or not
	var checkData model.User

	result := database.DB.Where("username = ?", user.Username).First(&checkData)
	log.Println("check data:", checkData)

	if result.Error != nil {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	} else if result.RowsAffected < 1 {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(checkData.Password), []byte(user.Password))
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	}

	// set JWT
	token, err := helpers.JWTGenerate(user.Username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token, // jwt token
		HttpOnly: true,
		Secure:   true,
	})

	// return 200
	model.Response(ctx, http.StatusOK, "Success")
}
