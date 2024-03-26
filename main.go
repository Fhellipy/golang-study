package main

import (
	"golang-study/api/routes"
	"golang-study/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Arquivo .env não encontrado")
	}

	app := fiber.New()

	// Routers config
	routes.SetupUserRoutes(app)

	db.InitDB()

	// Start the HTTP server on port 5000
	if err := app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
