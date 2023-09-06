package user_update_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_rest_api/internal/model"
	"golang_rest_api/internal/util/db"
	"net/http"
)

type RequestData struct {
	Id       int     `form:"id" json:"id" binding:"required"`
	Name     string  `form:"name" json:"name" binding:"required"`
	Email    string  `form:"email" json:"email" binding:"required,email"`
	Password *string `form:"password" json:"password" binding:"omitempty,max=255,min=8"`
}

var requestData RequestData
var user model.User

func validate(ctx *gin.Context) bool {

	err := ctx.ShouldBind(&requestData)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Valid error",
				"message": "Invalid inputs",
			})

		return false
	}

	return true
}

func Handle(ctx *gin.Context) {
	if validate(ctx) == false {
		return
	}

	dbConnection, _ := db.GetInstance()
	dbConnection.First(&user, requestData.Id)

	user.Name = requestData.Name
	user.Email = requestData.Email

	if requestData.Password != nil {
		password, _ := bcrypt.GenerateFromPassword([]byte(*requestData.Password), 14)
		user.Password = string(password)
	}

	dbConnection.Save(user)

	ctx.JSON(200, gin.H{
		"status": true,
	})
}
