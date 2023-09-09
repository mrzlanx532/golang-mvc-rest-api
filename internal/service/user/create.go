package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang_rest_api/internal/model"
	"golang_rest_api/internal/util/db"
	"net/http"
)

func NewCreateService() *CreateService {
	return &CreateService{}
}

type createServiceRequestData struct {
	Name     string `form:"name" json:"name" binding:"required,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,max=255,min=8"`
}

type CreateService struct {
}

var createServiceRequestDataLocal createServiceRequestData

func (cs CreateService) validate(ctx *gin.Context) bool {

	err := ctx.ShouldBind(&createServiceRequestDataLocal)

	if err != nil {

		fmt.Println(err)
		ctx.AbortWithStatusJSON(
			http.StatusLocked,
			gin.H{
				"errors":  "Validation error",
				"message": err.Error(),
			})

		return false
	}

	return true
}

func (cs CreateService) Handle(ctx *gin.Context) {

	if cs.validate(ctx) == false {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(createServiceRequestDataLocal.Password), 14)

	dbConnection, _ := db.GetInstance()
	dbConnection.Create(&model.User{
		Name:     createServiceRequestDataLocal.Name,
		Email:    createServiceRequestDataLocal.Email,
		Password: string(password),
	})

	ctx.JSON(200, gin.H{
		"status": true,
	})
}
