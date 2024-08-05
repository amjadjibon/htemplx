package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"htemplx/app/dto"
	"htemplx/app/models"
	"htemplx/app/repo"

	"github.com/google/uuid"
)

type ApiHandler struct {
	userRepo *repo.UsersRepo
}

func NewApiHandler(userRepo *repo.UsersRepo) ApiHandler {
	return ApiHandler{userRepo: userRepo}
}

// CreateUser example
//
//	@Summary	Add a new user
//	@Tags		Users
//	@Accept		json
//	@Produce	json
//	@Param		json	body		dto.CreateUserRequest	true	"Create user request"
//	@Success	201		{object}	dto.CreateUserResponse	"Create user response"
//	@Router		/users [post]
func (a *ApiHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
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

	err = a.userRepo.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	resp := dto.CreateUserResponse{
		ID: user.ID,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// GetUserList example
//
//	@Summary	Get all users
//	@Tags		Users
//	@Produce	json
//	@Success	200	{array}	dto.UserResponse	"User list response"
//	@Router		/users [get]
func (a *ApiHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users, err := a.userRepo.GetUserList(r.Context())
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	resp := make([]dto.UserResponse, len(users))
	for i, user := range users {
		resp[i] = dto.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// GetUserByID example
//
//	@Summary	Get a user
//	@Tags		Users
//	@Produce	json
//	@Param		id	path		string				true	"User ID"
//	@Success	200	{object}	dto.UserResponse	"User response"
//	@Router		/users/{id} [get]
func (a *ApiHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	user, err := a.userRepo.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	resp := dto.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// UpdateUser example
//
//	@Summary	Update user
//	@Tags		Users
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"User ID"
//	@Param		json	body		dto.UpdateUserRequest	true	"Update user request"
//	@Success	200		{object}	dto.UserResponse		"Update user response"
//	@Router		/users/{id} [put]
func (a *ApiHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := &models.User{
		ID:        uuid.MustParse(id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
	}

	err = a.userRepo.UpdateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	resp := dto.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// DeleteUser example
//
//	@Summary	Delete user
//	@Tags		Users
//	@Param		id	path	string	true	"User ID"
//	@Success	204	"User deleted successfully"
//	@Router		/users/{id} [delete]
func (a *ApiHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	err := a.userRepo.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
