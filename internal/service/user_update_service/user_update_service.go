package user_update_service

import (
	"fmt"
	//"golang_rest_api/internal/util/db"
	"net/http"
	"github.com/gin-gonic/gin"
	//"golang_rest_api/internal/model"
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

func Handle(ctx *gin.Context) {

	// TODO: to-do-do
}