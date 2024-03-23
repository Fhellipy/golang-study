package services

import (
	"golang-study/api/models"
	"golang-study/db"

	"github.com/gofiber/fiber/v2"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx *fiber.Ctx, user *models.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}
	collection := db.Database.Collection("user")

	if _, err := collection.InsertOne(ctx.Context(), user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) validateUser(user *models.User) error {
	// if user.Name == "" {
	// 	return models.ErrNameIsRequired
	// }

	// if user.Email == "" {
	// 	return models.ErrEmailIsRequired
	// }

	return nil
}
