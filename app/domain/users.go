package domain

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"htemplx/app/dto"
	"htemplx/app/models"
	"htemplx/app/repo"
	"htemplx/pkg/auth"
	"htemplx/pkg/mailer"
)

type UsersDomain struct {
	usersRepo *repo.UsersRepo
	mailer    *mailer.Mailer
}

func NewUsersDomain(
	usersRepo *repo.UsersRepo,
	mailer *mailer.Mailer,
) *UsersDomain {
	return &UsersDomain{usersRepo: usersRepo, mailer: mailer}
}

// CreateUsers processes the user creation request
func (u *UsersDomain) CreateUsers(r *http.Request) (*dto.CreateUserResponse, error) {
	var req dto.CreateUserRequest
	var err error

	// Check if the request expects JSON
	if r.Header.Get("Content-Type") == "application/json" {
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return nil, err
		}
	} else if r.Header.Get("HX-Request") == "true" &&
		r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		// Parse form data
		if err = r.ParseForm(); err != nil {
			return nil, err
		}

		req.FirstName = r.FormValue("first_name")
		req.LastName = r.FormValue("last_name")
		req.Username = r.FormValue("username")
		req.Email = r.FormValue("email")
		req.Password = r.FormValue("password")
	} else {
		return nil, errors.New("unsupported content type")
	}

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("missing required fields")
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user object
	user := &models.User{
		ID:        uuid.New(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	// Store user in repository
	err = u.usersRepo.CreateUser(r.Context(), user)
	if err != nil {
		return nil, err
	}

	// Return response
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

// Login processes the login form submission
func (u *UsersDomain) Login(r *http.Request) (*dto.UserResponse, error) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	// Extract form values
	email := r.FormValue("email")
	pass := r.FormValue("password")

	if email == "" || pass == "" {
		return nil, errors.New("missing required fields")
	}

	// Get user by email
	user, err := u.usersRepo.GetUserByEmail(r.Context(), email)
	if err != nil {
		return nil, err
	}

	// Validate password
	if !auth.CheckPassword(pass, user.Password) {
		return nil, errors.New("password incorrect")
	}

	// Return user response
	return &dto.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}, nil
}

func (u *UsersDomain) ForgotPassword(r *http.Request) error {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Extract form values
	email := r.FormValue("email")

	if email == "" {
		return errors.New("missing required fields")
	}

	// Get user by email
	user, err := u.usersRepo.GetUserByEmail(r.Context(), email)
	if err != nil {
		return err
	}

	rawPassword := auth.GenerateRandomPassword()
	user.Password, _ = auth.HashPassword(rawPassword)

	err = u.usersRepo.UpdateUser(r.Context(), user)
	if err != nil {
		return err
	}

	go func() {
		err = u.mailer.SendEmail(user.Email, "password", rawPassword)
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to send email", "error", err)
		}
	}()

	slog.InfoContext(r.Context(), "password reset successful")
	return nil
}
