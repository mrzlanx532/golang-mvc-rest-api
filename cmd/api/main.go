package main

import (
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/list/user_list"
	_ "golang_rest_api/internal/service"
	"golang_rest_api/internal/service/user_create_service"
	"golang_rest_api/internal/service/user_delete_service"
	"golang_rest_api/internal/service/user_update_service"
	"os"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/users/list", func(ctx *gin.Context) {
		user_list.Handle(ctx)
	})

	router.POST("/api/users/create", func(ctx *gin.Context) {
		user_create_service.Handle(ctx)
	})

	router.POST("/api/users/update", func(ctx *gin.Context) {
		user_update_service.Handle(ctx)
	})

	router.POST("/api/users/delete", func(ctx *gin.Context) {
		user_delete_service.Handle(ctx)
	})

	err := router.Run(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
