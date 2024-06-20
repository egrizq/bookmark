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
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, "Invalid login credentials.")
		return
	}
	log.Println("response user:", user)

	// check username is exist or not
	var userFromDB model.User

	result := database.DB.Where("username = ?", user.Username).First(&userFromDB)
	log.Println("check username:", userFromDB.Username)

	if result.Error != nil {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	} else if result.RowsAffected < 1 {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "Username of Password is invalid")
		return
	}

	// set session
	session, err := helpers.STORE.Get(ctx.Request, "session")
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	session.Values["username"] = user.Username
	log.Println("set session for:", session.Values["username"])

	if err := session.Save(ctx.Request, ctx.Writer); err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// return 200
	model.Response(ctx, http.StatusOK, user.Username)
	log.Println("login success")
}
