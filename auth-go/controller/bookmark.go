package controller

import (
	"auth-go/database"
	"auth-go/helpers"
	"auth-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBookmark(ctx *gin.Context) {
	// bind json
	requestData, err := helpers.RequestBookmark(ctx)
	if err != nil {
		model.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// get username by session
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// get user_id by username
	userID, err := helpers.GetUserID(username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// check category_name & get category id
	categoryID, err := helpers.CheckCategoryAndGetCategoryID(requestData.Category)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// insert new bookmark
	err = helpers.InsertBookmarkToDatabase(requestData, userID, categoryID)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// response 201
	model.Response(ctx, http.StatusCreated, "success")
}

func GetBookmarkByCategory(ctx *gin.Context) {
	// get params
	category := ctx.Param("category")

	// get username by session
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// get category id
	categoryID, err := helpers.CheckCategoryAndUserID(category, username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// get category
	listBookmark, err := helpers.GetBookmark(categoryID, username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// response 200
	model.Response(ctx, http.StatusOK, listBookmark)
}

func NewCategory(ctx *gin.Context) {
	// bind json
	category, err := helpers.RequestCategory(ctx)
	if err != nil {
		model.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// get username by session
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// get username id
	userID, err := helpers.GetUserID(username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// check category is exist or not
	err = helpers.CheckCategoryExistOrNot(userID, category)
	if err != nil {
		model.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// insert new category
	save := model.FormatCategory(userID, category)
	if result := database.DB.Save(&save); result.Error != nil {
		model.Response(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}

	model.Response(ctx, http.StatusCreated, "success")
}
