package handlers

import (
	"encoding/json"
	"htemplx/app/dto"
	"htemplx/app/models"
	"htemplx/app/repo"
	"net/http"
	"time"

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
//	@Accept		json
//	@Produce	json
//	@Param		json	body		dto.CreateUserRequest	true "Create user request"
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
	json.NewEncoder(w).Encode(resp)
}
