package main

import (
	"golang-study/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Routers config
	routes.SetupUserRoutes(app)

	// Start the HTTP server on port 5000
	if err := app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
