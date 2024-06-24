package helpers

import (
	"auth-go/database"
	"auth-go/model"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequestBookmark(ctx *gin.Context) (model.RequestBookmark, error) {
	var requestData model.RequestBookmark

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		return model.RequestBookmark{}, fmt.Errorf("invalid request")
	}
	log.Println("body bookmark:", requestData)

	return requestData, nil
}

func GetUserID(username string) (int, error) {
	var userID model.User
	err := database.DB.Select("id").Where("username = ?", username).First(&userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("user not found")
	} else if err != nil {
		log.Println("database error:", err.Error())
		return 0, fmt.Errorf("internal server error")
	}
	log.Println("user_id:", userID.ID)

	return int(userID.ID), nil
}

func CheckCategoryAndGetCategoryID(category string) (int, error) {
	var categoryID model.CategoryBookmark

	err := database.DB.Select("id").Where("category_name = ?", category).First(&categoryID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("category err:", err.Error())
		return 0, fmt.Errorf("category is not found")
	} else if err != nil {
		log.Println("database error:", err.Error())
		return 0, fmt.Errorf("internal server error")
	}
	log.Println("category_id:", categoryID.ID)

	return int(categoryID.ID), nil

}

func InsertBookmarkToDatabase(requestData model.RequestBookmark, userID int, categoryID int) error {
	newBookmark := model.FormatBookmark(requestData, userID, categoryID)
	log.Println("new bookmark:", newBookmark)

	result := database.DB.Save(&newBookmark)
	if result.Error != nil {
		log.Println("Database save error:", result.Error)
		return fmt.Errorf("internal server error")
	} else if result.RowsAffected < 1 {
		log.Println("No rows affected on save")
		return fmt.Errorf("internal server error")
	}

	return nil
}

func CheckCategoryAndUserID(category, username string) (int, error) {
	var categoryID int

	database.DB.Raw(`
		SELECT c.id
		FROM category_bookmarks c
		JOIN users u
		ON u.id = c.user_id
		WHERE u.username = ? AND c.category_name = ?
	`, username, category).Scan(&categoryID)
	if categoryID < 1 {
		return 0, fmt.Errorf("category is not found")
	}
	log.Println("id", categoryID)

	return categoryID, nil
}

func GetBookmark(categoryID int, username string) ([]model.ResponseBookmark, error) {
	var listBookmark []model.ResponseBookmark

	rows := database.DB.Raw(`
		SELECT b.social, b.url
		FROM bookmarks b
		JOIN users u
		ON u.id = b.user_id
		WHERE b.category_id = ? AND u.username = ?
	`, categoryID, username).Scan(&listBookmark)

	if errors.Is(rows.Error, gorm.ErrRecordNotFound) {
		log.Println("database error:", rows.Error.Error())
		return []model.ResponseBookmark{}, fmt.Errorf("please insert your bookmark")
	} else if rows.Error != nil {
		log.Println("database error:", rows.Error.Error())
		return []model.ResponseBookmark{}, fmt.Errorf("internal server error")
	}
	log.Println("list bookmark", listBookmark)

	return listBookmark, nil
}

func RequestCategory(ctx *gin.Context) (string, error) {
	var newCategory model.RequestCategory

	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		return "", fmt.Errorf("invalid request")
	}
	log.Println("body category:", newCategory)

	return newCategory.NewCategory, nil
}

func CheckCategoryExistOrNot(userID int, category string) error {
	var isCategoryExist model.CategoryBookmark

	database.DB.Raw(`
		SELECT c.category_name
		FROM category_bookmarks c
		WHERE c.user_id = ? AND c.category_name = ?
	`, userID, category).Scan(&isCategoryExist)

	if len(isCategoryExist.CategoryName) > 1 {
		log.Println("category name:", isCategoryExist.CategoryName)
		return fmt.Errorf("category already been used")
	}
	log.Println("category is not exist")

	return nil
}
