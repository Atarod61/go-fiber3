package db

import (
	"log"
	"os"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func Connect() {
	var err error
	// Connect to Database
	//dsn := os.Getenv("DB_DSN")
	dsn := "root:21@tcp(localhost:3308)/fiber?charset=utf8mb4&parseTime=True&loc=Local"
	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Get the underlying SQL database connection and defer its closure
	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()

	if err != nil {
		CloseDB()
		log.Fatal("error getting database connection: ", err)
	}
	log.Println("Database connected")

}
func AutoMigrate() {
	dbClient.AutoMigrate(&models.user{})
}

func GetDB() *gorm.DB {
	return dbClient
}
func CloseDB() {
	sqlDB, _ := dbClient.DB()
	sqlDB.Close()
}
file, err := os.Open("file.go") // For read access.
if err != nil {
    log.Fatal(err)
}