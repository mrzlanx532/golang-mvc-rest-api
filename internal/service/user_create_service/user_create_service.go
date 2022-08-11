package user_create_service

import (
	"golang_rest_api/internal/util/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func validate(ctx *gin.Context) {
	// TODO: Add validation to here..
}

func Execute(ctx *gin.Context) {

	validate(ctx)

	db, _ := db.GetConnection()
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Denis"})
}