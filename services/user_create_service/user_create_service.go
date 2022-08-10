package user_create_service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"os"
  )

type User struct {
	gorm.Model
	Name string
}

func Execute() {
	 dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	 if err != nil {
		fmt.Println(err)
		return
	 }

	 db.AutoMigrate(&User{})
	 db.Create(&User{Name: "Denis"})
}