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

	// get user id by category id
	userIdAndCategoryId, err := helpers.GetUserIDAndCategoryID(username, requestData.Category)
	if err != nil {
		model.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// insert new bookmark
	err = helpers.InsertBookmarkToDatabase(requestData,
		userIdAndCategoryId.UserID, userIdAndCategoryId.CategoryID)
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
		model.Response(ctx, http.StatusNotFound, err.Error())
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
		model.Response(ctx, http.StatusNoContent, err.Error())
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

func DeleteBookmark(ctx *gin.Context) {
	// get params
	id := ctx.Param("id")
	usernameParams := ctx.Param("username")
	log.Println("delete params:", usernameParams, id)

	// get session username
	username, err := helpers.GetSessionUsername(ctx)
	if err != nil {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// check username w usernameParams
	if username != usernameParams {
		model.Response(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// delete rows
	rows := database.DB.Exec(`
		DELETE
		FROM bookmarks
		WHERE id = ? AND user_id = (
			SELECT id 
			FROM users
			WHERE username = ?
		)
	`, id, usernameParams)

	if rows.Error != nil {
		log.Println("delete rows err:", rows.Error)
		model.Response(ctx, http.StatusInternalServerError, "internal server error")
		return
	} else if rows.RowsAffected == 0 {
		model.Response(ctx, http.StatusInternalServerError, "no rows affected")
		return
	}

	// response 200
	model.Response(ctx, http.StatusOK, "rows deleted")
}
