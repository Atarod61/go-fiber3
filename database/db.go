package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func Connect() {
	// Connect to Database
	//dsn := os.Getenv("DB_DSN")
	dsn := "root:21@tcp(localhost:3308)/fiber?charset=utf8mb4&parseTime=True&loc=Local"
	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Database connected")

}
