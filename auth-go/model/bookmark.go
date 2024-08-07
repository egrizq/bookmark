package model

import "gorm.io/gorm"

type RequestBookmark struct {
	Social   string `json:"social"`
	Url      string `json:"url"`
	Category string `json:"category"`
}

type Bookmark struct {
	gorm.Model
	UserID     int `gorm:"not null"`
	CategoryID int `gorm:"not null"`
	Social     string
	Url        string
	User       User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category   CategoryBookmark `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID"`
}

func FormatBookmark(request RequestBookmark, userID int, categoryId int) *Bookmark {
	return &Bookmark{
		UserID:     userID,
		Social:     request.Social,
		Url:        request.Url,
		CategoryID: categoryId,
	}
}

type ResponseBookmark struct {
	Id     string
	Url    string
	Social string
}

type RequestCategory struct {
	NewCategory string `json:"newcategory"`
}

type CategoryBookmark struct {
	gorm.Model
	UserID       int `gorm:"not null"`
	CategoryName string
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func FormatCategory(userId int, categoryName string) *CategoryBookmark {
	return &CategoryBookmark{
		UserID:       userId,
		CategoryName: categoryName,
	}
}

type CategoryAndBookmarksNumber struct {
	CategoryName string
	Number       int
}

type UserIDCategoryID struct {
	UserID     int
	CategoryID int
}

type CategoryAndDateTime struct {
	Hour         string
	Date         string
	CategoryName string
}
