package config

import (
	"fmt"
	"go-native-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE, "%2F")

	db, error := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if error != nil {
		panic("Failed to connect database")
	}


	db.AutoMigrate(&models.Author{})

	DB = db
	log.Println("Database connected")
}