package user

import (
	"golang-study/api/models"
	"golang-study/api/services"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	us := services.NewUserService()

	if err := us.CreateUser(ctx, &user); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuário cadastrado com sucesso",
		"user":    user,
	})
}
