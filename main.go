package main

import (
	"fmt"
	"log"
	"os"

	db "go-fiber3/database"
	"go-fiber3/models"

	//"go-fiber3/seeders"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found. Defaulting to environment variables.")
	}

	//DB stuff
	db.Connect()
	db.AutoMigrate()
	//seeders.SeedUsers()

	// Set up a simple GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Everything is working fine",
		})
	})
	app.Get("/add-user", func(c *fiber.Ctx) error {
		user := models.User{
			Username: "super-admin",
			Password: "superadmin@123",
			Email:    "superadmin@gmail.com",
		}
		dbclient := db.GetDB()
		return dbclient.Create(&user).Error
	})
	//	app.Get("/delet-user/:id",func(c *fiber.Ctx) error {
	//id : c.Params("id")
	//userId,- := strerrconv.Atoi(id)
	// user := models.User{
	//	Model: gorm.Model{
	//	ID: uint(userId),
	//},
	//	 }
	// dbclient := db.GetDB()
	//dbclient.Model(&user).Tr(
	//	err := dbclient.Delete(&user).Error
	// return err
	//)
	//})

	// Start the Fiber app
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Default to port 3000 if PORT is not set
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
