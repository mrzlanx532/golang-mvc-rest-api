package service

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang_rest_api/internal/util/db"
	"os"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Не найден .env файл")
	}

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	_, e := db.GetInstance()

	if e != nil {
		panic(e)
	}
}
