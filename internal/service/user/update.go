package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_rest_api/internal/model"
	"golang_rest_api/internal/util/db"
	"net/http"
)

func NewUpdateService() *UpdateService {
	return &UpdateService{}
}

type UpdateService struct {
	updateServiceRequestData
}

type updateServiceRequestData struct {
	Id       int     `form:"id" json:"id" binding:"required"`
	Name     string  `form:"name" json:"name" binding:"required"`
	Email    string  `form:"email" json:"email" binding:"required,email"`
	Password *string `form:"password" json:"password" binding:"omitempty,max=255,min=8"`
}

var updateServiceRequestDataLocal updateServiceRequestData
var user model.User

func (us UpdateService) validate(ctx *gin.Context) bool {

	err := ctx.ShouldBind(&updateServiceRequestDataLocal)

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

func (us UpdateService) Handle(ctx *gin.Context) {

	userUpdateService := NewUpdateService()

	if userUpdateService.validate(ctx) == false {
		return
	}

	dbConnection, _ := db.GetInstance()
	dbConnection.First(&user, updateServiceRequestDataLocal.Id)

	user.Name = updateServiceRequestDataLocal.Name
	user.Email = updateServiceRequestDataLocal.Email

	if updateServiceRequestDataLocal.Password != nil {
		password, _ := bcrypt.GenerateFromPassword([]byte(*updateServiceRequestDataLocal.Password), 14)
		user.Password = string(password)
	}

	dbConnection.Save(user)

	ctx.JSON(200, gin.H{
		"status": true,
	})
}
