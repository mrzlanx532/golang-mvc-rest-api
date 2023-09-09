package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/model"
	"golang_rest_api/internal/util/db"
	"net/http"
)

func NewDeleteService() *DeleteService {
	return &DeleteService{}
}

type DeleteService struct {
}

type deleteServiceRequestData struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

var deleteServiceRequestDataLocal deleteServiceRequestData

func (ds DeleteService) validate(ctx *gin.Context) bool {

	err := ctx.ShouldBind(&deleteServiceRequestDataLocal)

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

func (ds DeleteService) Handle(ctx *gin.Context) {

	if ds.validate(ctx) == false {
		return
	}

	dbConnection, _ := db.GetInstance()
	dbConnection.Delete(&model.User{}, deleteServiceRequestDataLocal.ID)

	ctx.JSON(200, gin.H{
		"status": true,
	})
}
