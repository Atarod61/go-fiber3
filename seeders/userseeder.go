package seeders

import (
	"log"

	db "go-fiber3/database"
	"go-fiber3/models"
)

func SeedUsers() {

	dbClient := db.GetDB()
	// ravesh 1
	//var user models.use
	//user.Email = "test@test.com"
	//user.Username = "test"
	//user.Password = "123456789"

	//ravesh 2
	User := models.user{
		Username : "Atarod61",
		Password : "123456789",
		Email : "atarod2017@gmail.com"
	}

	err := dbClient.Create(&user).Error

	if err != nil {
		log.Fatalf("user cannot be added to database")
	}

}
