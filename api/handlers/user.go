package handlers

import (
	"golang-study/api/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(app *fiber.App) {
	userRouter := app.Group("/api/users")
	userRouter.Post("/", user.CreateUser)
}
