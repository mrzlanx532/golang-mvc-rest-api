package db

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConnectionInstance *gorm.DB

func GetConnection() (*gorm.DB, error) {
	if dbConnectionInstance == nil {

		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
		dbConnectionInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	 	if err != nil {
			fmt.Println(err)
			log.Fatal("Пиздец")
			
			return nil, err
	 	}

		return dbConnectionInstance, nil
	}

	log.Println("Соединение уже установлено")

	return dbConnectionInstance, nil
}