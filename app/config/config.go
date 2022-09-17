package config

import (
	"fmt"

	"github.com/adityaerlangga/golang-auth/entities/userentity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DB_HOST := "localhost"
	DB_NAME := "golang_auth"
	DB_USER := "root"
	DB_PASS := ""
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME)
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&userentity.User{})
	DB = database
}
