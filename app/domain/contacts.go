package domain

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"htemplx/app/dto"
	"htemplx/app/models"
	"htemplx/app/repo"
)

type ContactsDomain struct {
	contactsRepo *repo.ContactsRepo
}

func NewContactsDomain(contactsRepo *repo.ContactsRepo) *ContactsDomain {
	return &ContactsDomain{contactsRepo: contactsRepo}
}

// CreateContacts processes the user creation request
func (u *ContactsDomain) CreateContacts(r *http.Request) error {
	var req dto.CreateContactsRequest
	var err error

	// Check if the request expects JSON
	if r.Header.Get("Content-Type") == "application/json" {
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}
	} else if r.Header.Get("HX-Request") == "true" &&
		r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		// Parse form data
		if err = r.ParseForm(); err != nil {
			return err
		}

		req.Email = r.FormValue("email")
		req.Subject = r.FormValue("subject")
		req.Message = r.FormValue("message")
	} else {
		return errors.New("unsupported content type")
	}

	if req.Email == "" || req.Subject == "" || req.Message == "" {
		return errors.New("missing required fields")
	}

	// Create user object
	contacts := &models.ContactUs{
		Email:     req.Email,
		Subject:   req.Subject,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}

	// Store user in repository
	err = u.contactsRepo.CreateContacts(r.Context(), contacts)
	if err != nil {
		return err
	}

	return nil
}
