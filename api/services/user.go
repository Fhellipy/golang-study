package services

import (
	"fmt"
	"golang-study/api/models"
	"golang-study/api/utils"
	"golang-study/db"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx *fiber.Ctx, user *models.UserCreate) error {
	collection := db.Database.Collection("user")

	if err := s.validateUser(ctx, collection, user); err != nil {
		return err
	}

	user.ID = uuid.New().String()
	timeNow := time.Now()
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow

	// Hash password
	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.RequestErrorValidation(ctx, err)
	}

	user.Password = passwordHashed

	fmt.Println("collection", collection)

	if _, err := collection.InsertOne(ctx.Context(), user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) validateUser(ctx *fiber.Ctx, collection *mongo.Collection, user *models.UserCreate) error {
	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return err
	}

	filter := bson.M{"email": user.Email}

	if err := collection.FindOne(ctx.Context(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			erro := models.CustomError{
				Key:   "bad_request",
				Value: "email já cadastrado",
			}

			return &erro
		}

		erro := models.Error{
			Message: "Erro ao buscar usuário",
			Err:     "internal_server_error",
			Code:    500,
		}

		return &erro
	}

	return nil
}
