package controller

import (
	"auth-go/database"
	"auth-go/helpers"
	"auth-go/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBookmark(ctx *gin.Context) {
	// bind json
	var requestData model.RequestBookmark

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		model.Response(ctx, http.StatusBadRequest, "invalid request")
		return
	}
	log.Println(requestData)

	// get id by username
	var id model.User
	result := database.DB.Select("id").Where("username = ?", requestData.Username).Find(&id)
	if result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	} else if result.RowsAffected < 1 {
		model.Response(ctx, http.StatusBadRequest, "username is not found")
		return
	}

	// insert bookmark
	save := model.FormatSaveBookmark(requestData, int(id.ID))
	if result := database.DB.Create(save); result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}

	// response 201
	model.Response(ctx, http.StatusCreated, "success")
}

func GetBookmarkByCategory(ctx *gin.Context) {
	// get params
	params := ctx.Param("category")

	// get username by session
	session, err := helpers.STORE.Get(ctx.Request, "session")
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, "cannot get session")
		return
	}

	username, ok := session.Values["username"].(string)
	if !ok {
		model.Response(ctx, http.StatusInternalServerError, "cannot get session")
		return
	}

	// check category is exist or not
	// db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

	type res struct {
		Url    string
		Social string
	}

	var get []res
	database.DB.Table("bookmarks").Select("bookmarks.social, bookmarks.url").Joins("join users on users.id = bookmarks.user_id").Where("users.username = ? AND bookmarks.category = ?", username, params).Scan(&get)

	log.Println(get)
	// response 200
	model.Response(ctx, http.StatusOK, get)
}

func NewCategory(ctx *gin.Context) {

}
