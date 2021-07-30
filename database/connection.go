package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {

	_, err := gorm.Open(mysql.Open("dev:123456@/go_jwt"), &gorm.Config{})
	if err != nil {
		panic("couldn't establish database connection")
	}

}