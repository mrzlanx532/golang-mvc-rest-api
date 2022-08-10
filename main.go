package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang_rest_api/services/user_create_service"
)

func main() {

	godotenv.Load()

	router := gin.Default()
	
	router.GET("/ping", func(c *gin.Context) {
		user_create_service.Execute()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/users/create", func(ctx *gin.Context) {
		user_create_service.Execute()
	})

	fmt.Println("test")

	router.Run(":"+os.Getenv("PORT"))
}
