package user_create_service

import (
	"fmt"
	"golang_rest_api/internal/util/db"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/model"
)

type RequestData struct {
	Name string `form:"name" json:"name" binding:"required"`
}

var requestData RequestData

func validate(ctx *gin.Context) {
	
	fmt.Print("hello")
	err := ctx.ShouldBind(&requestData)

	if err!= nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"error": "Valid error",
			"message": "Invalid inputs",
		})
	}
}

func Execute(ctx *gin.Context) {

	validate(ctx)

	db, _ := db.GetConnection()
	db.AutoMigrate(&model.User{})
	db.Create(&model.User{Name: requestData.Name})

	ctx.JSON(200, gin.H{
		"status": true,
	})
}