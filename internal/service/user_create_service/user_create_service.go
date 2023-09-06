package user_create_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/model"
	"golang_rest_api/internal/util/db"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type ApiError struct {
	Field string
	Message string
}

type RequestData struct {
	Name string `form:"name" json:"name" binding:"required,max=255"`
	Email string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,max=255,min=8"`
}

var requestData RequestData

func validate(ctx *gin.Context) bool {
	
	err := ctx.ShouldBind(&requestData)

	if err != nil {

		fmt.Println(err)
		ctx.AbortWithStatusJSON(
			http.StatusLocked,
			gin.H{
				"errors": "Validation error",
				"message": err.Error(),
			})

		return false
	}

	return true
}

func Handle(ctx *gin.Context) {

	if validate(ctx) == false {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(requestData.Password), 14)

	dbConnection, _ := db.GetInstance()
	dbConnection.Create(&model.User{
		Name: requestData.Name,
		Email: requestData.Email,
		Password: string(password),
	})

	ctx.JSON(200, gin.H{
		"status": true,
	})
}