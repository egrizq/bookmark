package helpers

import (
	"auth-go/database"
	"auth-go/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func BindJSON(ctx *gin.Context) (model.RequestUser, error) {
	var user model.RequestUser

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return model.RequestUser{}, fmt.Errorf("invalid login credentials")
	}
	log.Println("body user:", user)

	return user, nil
}

func IsUsernameExistForLogin(username string) (string, error) {
	var userFromDB model.User

	result := database.DB.Where("username = ?", username).First(&userFromDB)
	if result.Error != nil {
		return "", result.Error
	} else if result.RowsAffected < 1 {
		return "", result.Error
	}
	log.Println("check username is success")

	return userFromDB.Password, nil
}

func IsUsernameExistForSignUp(username string) error {
	var userFromDB model.User

	database.DB.Where("username = ?", username).First(&userFromDB)
	if userFromDB.Username != "" {
		return fmt.Errorf("username already exist")
	}
	log.Println("check username is success")

	return nil
}

func SetSession(ctx *gin.Context, username string) error {
	session, err := STORE.Get(ctx.Request, "session")
	if err != nil {
		return err
	}
	session.Values["username"] = username

	if err := session.Save(ctx.Request, ctx.Writer); err != nil {
		return err
	}
	log.Println("session is set for:", username)

	return nil
}

func GetSessionUsername(ctx *gin.Context) (string, error) {
	session, err := STORE.Get(ctx.Request, "session")
	if err != nil {
		return "", err
	}

	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		return "", err
	}
	log.Println("username:", username)

	return username, nil
}
