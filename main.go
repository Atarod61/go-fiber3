package main

import(
"log"  
"os"
"github.com/gofiber/fiber/v2"
"github.com/joho/godotenv"
 "gorm.io/driver/mysql"
 "gorm.io/gorm"
 
var dbclient*gorm.DB
func main(){
	//Initialize.fiber
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file")
	}
	//connnect to DB
	dsn := os.Getven("DB-DSN")

	dbliient, err  =gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	sqldb, err := dbclient.DB()
	defer sqldb close()
	if err != nil {
		log.Fatal("error connecting to database ")
	}
	log.Println("Database connected")

app.Get("/",func(c*fiber.ctx)) error{
	return c.status(fiber.statusok).json(&fiber.Map{
		"message" : "Everything is working fin",
	})
})
app.Listen(fmt.sprintf(":%s",os.Getven("PORT")))