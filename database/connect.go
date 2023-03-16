package database

import (
	"go-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("example-user:my_cool_secret@/go-admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connet to the database")
	}

	DB = database
	database.AutoMigrate(&models.User{})
}
