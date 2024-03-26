package models

import (
	"fmt"
	"net/http"
)

type Error struct {
	// Error message.
	Message string `json:"message" bson:"message"`

	// Error description.
	Err string `json:"error" bson:"error"`

	// Error code.
	Code int `json:"code" bson:"code"`

	// Field errors.
	FieldErrors map[string]string `json:"field_errors" bson:"field_errors"`
}

type CustomError struct {
	Key   string
	Value string
}

func (r *Error) Error() string {
	return r.Message
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Key, e.Value)
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
