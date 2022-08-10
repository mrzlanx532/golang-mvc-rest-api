package user_create_service

import (
	"gorm.io/gorm"
	"golang_rest_api/internal/util/db"
)

type User struct {
	gorm.Model
	Name string
}

func Execute() {

	 db, _ := db.GetConnection()

	 db.AutoMigrate(&User{})
	 db.Create(&User{Name: "Denis"})
}