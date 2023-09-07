package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var dbConnectionInstance *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if dbConnectionInstance == nil {

		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
		dbConnectionInstanceLocal, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			return nil, err
		}

		dbConfiguration, err := dbConnectionInstanceLocal.DB()

		if err != nil {
			fmt.Println(err)
			log.Fatal("Не удалось установить соединение с mysql")

			return nil, err
		}

		dbConfiguration.SetMaxIdleConns(10)
		dbConfiguration.SetMaxOpenConns(100)
		dbConfiguration.SetConnMaxLifetime(60 * 5)

		dbConnectionInstance = dbConnectionInstanceLocal

		return dbConnectionInstance, nil
	}

	return dbConnectionInstance, nil
}
