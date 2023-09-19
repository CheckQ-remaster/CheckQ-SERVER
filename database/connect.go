package database

import (
	"log"

	"example.com/m/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {

	USER := "root"
	PASS := "1234"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "CHECKQ"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("gorm DB Open Error")
		panic(err)
	}

	if err = db.AutoMigrate(&models.User{}, &models.Hotel{}); err != nil {
		panic(err)
	}

	DB = db
}
