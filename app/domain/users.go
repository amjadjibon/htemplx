package domain

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"htemplx/app/dto"
	"htemplx/app/models"
	"htemplx/app/repo"
)

type UsersDomain struct {
	usersRepo *repo.UsersRepo
}

func NewUsersDomain(usersRepo *repo.UsersRepo) *UsersDomain {
	return &UsersDomain{usersRepo: usersRepo}
}

func (u *UsersDomain) CreateUsers(r *http.Request) (*dto.CreateUserResponse, error) {
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("missing required fields")
	}

	user := &models.User{
		ID:        uuid.New(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	err = u.usersRepo.CreateUser(r.Context(), user)
	if err != nil {
		return nil, err
	}

	resp := &dto.CreateUserResponse{
		ID: user.ID,
	}

	return resp, nil
}

func (u *UsersDomain) GetUserList(r *http.Request) ([]dto.UserResponse, error) {
	users, err := u.usersRepo.GetUserList(r.Context())
	if err != nil {
		return nil, err
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
		})
	}

	return userResponses, nil
}

func (u *UsersDomain) GetUserByID(r *http.Request) (*dto.UserResponse, error) {
	idParam := chi.URLParam(r, "id")
	user, err := u.usersRepo.GetUserByID(r.Context(), idParam)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	resp := &dto.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	return resp, nil
}

func (u *UsersDomain) UpdateUser(r *http.Request) (*dto.UserResponse, error) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return nil, err
	}

	var req dto.UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		UpdatedAt: sql.Null[time.Time]{Valid: true, V: time.Now()},
	}

	err = u.usersRepo.UpdateUser(r.Context(), user)
	if err != nil {
		return nil, err
	}

	resp := &dto.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	return resp, nil
}

func (u *UsersDomain) DeleteUser(r *http.Request) (*dto.DeleteUserResponse, error) {
	idParam := chi.URLParam(r, "id")
	err := u.usersRepo.DeleteUser(r.Context(), idParam)
	if err != nil {
		return nil, err
	}

	resp := &dto.DeleteUserResponse{
		Message: "User deleted successfully",
	}

	return resp, nil
}
