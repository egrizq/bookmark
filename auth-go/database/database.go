package database

import (
	"auth-go/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	db := os.Getenv("DB")
	password := os.Getenv("PASSWORD")

	config := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		host, port, user, db, password)

	connect, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	if err := connect.AutoMigrate(&model.User{}, &model.Bookmark{}, &model.CategoryBookmark{}); err != nil {
		log.Fatal("Error migrate model", err)
	}

	DB = connect

	log.Println("Database is connected")
}
