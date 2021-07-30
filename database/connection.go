package database

import (
	"github.com/cryptus-neoxys/go-spa-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	conn, err := gorm.Open(mysql.Open("dev:123456@/go_jwt"), &gorm.Config{})
	if err != nil {
		panic("couldn't establish database connection")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}