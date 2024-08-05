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
