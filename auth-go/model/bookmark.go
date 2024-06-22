package model

import "gorm.io/gorm"

type RequestBookmark struct {
	Username string `json:"username"`
	Social   string `json:"social"`
	Url      string `json:"url"`
	Category string `json:"category"`
}

type Bookmark struct {
	gorm.Model
	UserID   int `gorm:"not null"`
	Social   string
	Url      string
	Category string
	User     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func FormatSaveBookmark(request RequestBookmark, id int) *Bookmark {
	return &Bookmark{
		UserID:   id,
		Social:   request.Social,
		Url:      request.Url,
		Category: request.Category,
	}
}
