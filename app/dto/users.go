package dto

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name" example:"Jon"`
	LastName  string `json:"last_name" example:"Doe"`
	Username  string `json:"username" example:"jon.doe"`
	Email     string `json:"email" example:"jon.doe@gmail.com"`
	Password  string `json:"password" example:"******"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id" example:"1879829d-0252-4d03-bd9d-980af59dfe2b"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id" example:"1879829d-0252-4d03-bd9d-980af59dfe2b"`
	FirstName string    `json:"first_name" example:"Jon"`
	LastName  string    `json:"last_name" example:"Doe"`
	Username  string    `json:"username" example:"jon.doe"`
	Email     string    `json:"email" example:"jon.doe@gmail.com"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" example:"Jon"`
	LastName  string `json:"last_name" example:"Doe"`
	Username  string `json:"username" example:"jon.doe"`
	Email     string `json:"email" example:"jon.doe@gmail.com"`
	Password  string `json:"password" example:"******"`
}

type DeleteUserResponse struct {
	Message string `json:"message" example:"User deleted successfully"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"An error occurred"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"jon.doe@gmail.com"`
	Password string `json:"password" example:"******"`
}
