package main

import (
	"log"
	"os"

	"github.com/Ibnuardhian/blogwebbackend/database"
	"github.com/Ibnuardhian/blogwebbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env Files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
