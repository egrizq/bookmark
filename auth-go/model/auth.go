package model

import "gorm.io/gorm"

type RequestUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Email    string
	Username string
	Password string
	Bookmark []Bookmark `gorm:"foreignKey:UserID"`
}

func FormatUser(request RequestUser) *User {
	return &User{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}
}
