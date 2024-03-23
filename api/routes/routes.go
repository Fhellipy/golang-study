package routes

import (
	"golang-study/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	handlers.CreateUserHandler(app)
}
