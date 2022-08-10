package main

import (
	"golang_rest_api/internal/service/user_create_service"
	"golang_rest_api/internal/util/db"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
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

	router.GET("/api/users/list", func(ctx *gin.Context) {
		db, _ := db.GetConnection()
		rows, _ := db.Raw("SELECT * FROM users").Rows()

		pp.Print(rows)

		for rows.Next() {
			defer rows.Close()
			// db.ScanRows(&name)
		}
	})

	router.POST("/api/users/create", func(ctx *gin.Context) {
		user_create_service.Execute()
	})

	router.Run(":"+os.Getenv("PORT"))
}
