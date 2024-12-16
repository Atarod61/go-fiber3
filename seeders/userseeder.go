package seeders

import (
	"log"

	db "github.com/Atarod61/go-fiber3/database"
)

func seedusers() {

	dbClient := db.GetDB()

	error := dbClient.create().Error

	if err != nil {
		log.Fatalf("user cannot be added to database")
	}

}
