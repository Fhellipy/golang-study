package models

import "time"

type UserRequest struct {
	Name     string `json:"name" bson:"name" validate:"required,min=3,max=60"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8,max=60,containsany=!@#$%&*"`
}
type UserCreate struct {
	ID        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name" validate:"required,min=3,max=60"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password" bson:"password" validate:"required,min=8,max=60,containsany=!@#$%&*"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
