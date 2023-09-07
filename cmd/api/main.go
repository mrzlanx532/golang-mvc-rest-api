package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/controller"
	_ "golang_rest_api/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userController := controller.UserController{}

	router.GET("/api/users/list", userController.List())
	router.POST("/api/users/create", userController.Create())
	router.POST("/api/users/update", userController.Update())
	router.POST("/api/users/delete", userController.Delete())

	err := router.Run(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

	gracefulShutdown()
}

func gracefulShutdown() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	_ = <-signalChannel

	fmt.Println("Завершение приложения")
}
