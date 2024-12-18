package seeders

import (
	"log"

	db "github.com/Atarod61/go-fiber3/database"
)

func seedusers() {

	dbClient := db.GetDB()
	// ravesh 1
	var user models.user
	user.Email = "test@test.com"
	user.Username = "test"
	user.Password = "123456789"

	//ravesh 2

	err := dbClient.Create(&user).Error

	if err != nil {
		log.Fatalf("user cannot be added to database")
	}

}
