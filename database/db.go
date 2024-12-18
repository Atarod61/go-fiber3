package db

import (
	"database/sql"
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

	// Get the underlying SQL database connection and defer its closure
	sqlDB,_ := dbClient.DB()
	err = sqlDB.Ping()

	 err != nil {
		CloseDB()
		log.Fatal("Error getting database connection: ", err)
	}
	log.Println("Database connected")

}
func AutoMigrate() {
	dbClient.AutoMigrate(&models.user{})
}

func GetDB() *gorm.DB {
	return dbClient
}
func CloseDB(){
	sqlDB,_ := dbClient.DB()
	sqlDB.Close()
}
