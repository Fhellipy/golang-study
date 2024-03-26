package utils

import (
	"golang-study/api/models"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func msgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "Campo obrigatório"
	case "email":
		return "Email inválido"
	case "min":
		return "O tamanho mínimo é " + param
	case "max":
		return "O tamanho máximo é " + param
	case "datetime":
		return "Data inválida"
	case "numeric":
		return "O valor deve ser numérico"
	case "containsany":
		return "O valor deve conter pelo menos um destes caracteres: " + param
	default:
		return "Campo inválido"
	}
}

func RequestErrorValidation(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case validator.ValidationErrors:
		validationErrors := make(map[string]string)

		for _, fieldError := range e {
			message := msgForTag(fieldError.Tag(), fieldError.Param())
			key := strings.ToLower(fieldError.Field())
			validationErrors[key] = message
		}

		errorResponse := struct {
			Message     string            `json:"message"`
			Err         string            `json:"err"`
			Code        int               `json:"code"`
			FieldErrors map[string]string `json:"field_errors"`
		}{
			Message:     "Alguns campos estão incorretos",
			Err:         "bad_request",
			Code:        400,
			FieldErrors: validationErrors,
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)

	case *models.CustomError:
		errorResponse := struct {
			Message string `json:"message"`
			Err     string `json:"err"`
			Code    int    `json:"code"`
		}{
			Message: e.Message,
			Err:     e.Err,
			Code:    400,
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)

	default:
		errorResponse := struct {
			Message string `json:"message"`
			Err     string `json:"err"`
			Code    int    `json:"code"`
		}{
			Message: "Erro interno do servidor",
			Err:     "internal_server_error",
			Code:    500,
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse)
	}
}
