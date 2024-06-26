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
	username := ctx.Param("username")
	category := ctx.Param("category")

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

func InsertNewCategory(ctx *gin.Context) {
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
	log.Println("new category", save.CategoryName)

	model.Response(ctx, http.StatusCreated, "success")
}

func GetListCategory(ctx *gin.Context) {
	// get session username
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// get list category by username
	listCategory, err := helpers.GetCategoryByUsername(username)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// response 200
	model.Response(ctx, http.StatusOK, listCategory)
}

func GetListOfCategoryAndNumberOfBookmarks(ctx *gin.Context) {
	// get session username
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// get category with number of bookmarks
	categoryAndNumberBookmarks, err := helpers.GetCategoryAndNumberOfBookmarks(username)
	if err != nil {
		model.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// response 200
	model.Response(ctx, http.StatusOK, categoryAndNumberBookmarks)
}
