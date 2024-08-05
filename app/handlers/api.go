package handlers

import (
	"encoding/json"
	"net/http"

	"htemplx/app/domain"
	"htemplx/app/repo"
)

type ApiHandler struct {
	usersDomain *domain.UsersDomain
}

func NewApiHandler(userRepo *repo.UsersRepo) ApiHandler {
	return ApiHandler{
		usersDomain: domain.NewUsersDomain(userRepo),
	}
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
	resp, err := a.usersDomain.CreateUsers(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetUserList example
//
//	@Summary	Get all users
//	@Tags		Users
//	@Produce	json
//	@Success	200	{array}	dto.UserResponse	"User list response"
//	@Router		/users [get]
func (a *ApiHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	resp, err := a.usersDomain.GetUserList(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	resp, err := a.usersDomain.GetUserByID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	resp, err := a.usersDomain.UpdateUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteUser example
//
//	@Summary	Delete user
//	@Tags		Users
//	@Param		id	path	string	true	"User ID"
//	@Success	204	"User deleted successfully"
//	@Router		/users/{id} [delete]
func (a *ApiHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := a.usersDomain.DeleteUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
