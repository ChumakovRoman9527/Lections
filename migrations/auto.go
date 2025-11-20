package main

import (
	"12-Context/internal/link"
	"12-Context/internal/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&link.Link{}, &user.User{})
	if err != nil {
		panic(err)
	}
}
