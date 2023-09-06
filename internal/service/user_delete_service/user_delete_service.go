package user_delete_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/model"
	"golang_rest_api/internal/util/db"
	"net/http"
)

type RequestData struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

var requestData RequestData

func validate(ctx *gin.Context) {

	err := ctx.ShouldBind(&requestData)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Valid error",
				"message": "Invalid inputs",
			})
	}
}

func Handle(ctx *gin.Context) {

	validate(ctx)

	dbConnection, _ := db.GetInstance()
	dbConnection.Delete(&model.User{}, requestData.ID)

	ctx.JSON(200, gin.H{
		"status": true,
	})
}
