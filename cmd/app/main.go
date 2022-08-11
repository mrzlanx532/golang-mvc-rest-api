package main

import (
	"golang_rest_api/internal/service/user_create_service"
	"golang_rest_api/internal/util/db"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//"github.com/k0kubun/pp"
)

func main() {

	godotenv.Load()

	router := gin.Default()
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/users/list", func(ctx *gin.Context) {
		db, _ := db.GetConnection()

		type Result struct {
			Id int64 `json:"id"`
			Name string `json:"name"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
			DeletedAt *string `json:"deleted_at"`
		}

		var result Result

		rows, _ := db.Raw("SELECT * FROM users").Rows()

		data := make([]Result, 0, 50)

		defer rows.Close()

		for i:=0; rows.Next(); i++ {
			db.ScanRows(rows, &result)
			data = append(data, result)
		}
		
		//fmt.Print("%+v", data) 
		//pp.Print(data)
		fmt.Println(len(data))
		fmt.Println(cap(data))

		ctx.JSON(200, data)
	})

	router.POST("/api/users/create", func(ctx *gin.Context) {
		user_create_service.Execute(ctx)
	})

	router.Run(":"+os.Getenv("PORT"))
}
