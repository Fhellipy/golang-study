package user

import (
	"golang-study/api/models"
	"golang-study/api/services"
	"golang-study/api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user models.UserCreate

	if err := ctx.BodyParser(&user); err != nil {
		error := models.NewBadRequestError("Alguns campos estão incorretos ou faltando")
		return ctx.Status(error.Code).JSON(error)
	}

	us := services.NewUserService()

	if err := us.CreateUser(ctx, &user); err != nil {
		return utils.RequestErrorValidation(ctx, err)
	}

	// remove password from response
	user.Password = ""

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuário cadastrado com sucesso",
		"user":    user,
	})
}
