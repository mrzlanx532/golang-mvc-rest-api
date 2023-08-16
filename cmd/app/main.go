package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang_rest_api/internal/service/user_create_service"
	"golang_rest_api/internal/service/user_delete_service"
	"golang_rest_api/internal/service/user_update_service"
	"golang_rest_api/internal/util/db"
	"os"
)

func main() {

	godotenv.Load()

	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	dbConnection, _ := db.GetConnection()

	router.GET("/api/users/list", func(ctx *gin.Context) {

		type Result struct {
			Id int64 `json:"id"`
			Name string `json:"name"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
			DeletedAt *string `json:"deleted_at"`
		}

		var result Result

		rows, _ := dbConnection.Raw("SELECT * FROM users WHERE deleted_at is NULL").Rows()

		data := make([]Result, 0, 50)

		defer rows.Close()

		for i:=0; rows.Next(); i++ {
			dbConnection.ScanRows(rows, &result)
			data = append(data, result)
		}
		
		ctx.JSON(200, data)
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

	router.Run(":"+os.Getenv("PORT"))
}
