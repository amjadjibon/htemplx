package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"

	"htemplx/app/domain"
	"htemplx/app/views/components"
	"htemplx/app/views/pages"
)

type WebHandler struct {
	usersDomain  *domain.UsersDomain
	sessionStore sessions.Store
}

func NewWebHandler(usersDomain *domain.UsersDomain, sessionStore sessions.Store) WebHandler {
	return WebHandler{usersDomain: usersDomain, sessionStore: sessionStore}
}

// Index renders the index page
func (h *WebHandler) Index(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := h.IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.Index(loggedIn, "htemplx"))
}

// About renders the about page
func (h *WebHandler) About(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := h.IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.About(loggedIn, "htemplx"))
}

// Contact renders the contact page
func (h *WebHandler) Contact(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := h.IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.Contact(loggedIn, "htemplx"))
}

// NotFound renders the 404 not found page
func (h *WebHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, r, pages.NotFound("htemplx"))
}

// Services renders the services page
func (h *WebHandler) Services(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := h.IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.Services(loggedIn, "htemplx"))
}

// Login renders the login page
func (h *WebHandler) Login(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.Login())
}

// SignIn renders the pastebin page (adjust the name if needed)
func (h *WebHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	_, err := h.usersDomain.Login(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := h.sessionStore.New(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.NavPasteBin(true))
}

// SignOut renders the pastebin page (adjust the name if needed)
func (h *WebHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	session, err := h.sessionStore.New(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = false
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Register renders the registration page
func (h *WebHandler) Register(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.Register())
}

// SignUp renders the registration page
func (h *WebHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	_, err := h.usersDomain.CreateUsers(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render(w, r, components.Login())
}

// ForgotPassword renders the forgot password page
func (h *WebHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.ForgotPassword())
}

// UnderConstruction renders the under construction page
func (h *WebHandler) UnderConstruction(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.UnderConstruction())
}

// TermsAndConditions renders the terms and conditions page
func (h *WebHandler) TermsAndConditions(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := h.IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.Terms(loggedIn, "htemplx"))
}

// Privacy renders the privacy policy page
func (h *WebHandler) Privacy(w http.ResponseWriter, r *http.Request) {
	loggedIn, err := h.IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render(w, r, pages.Privacy(loggedIn, "htemplx"))
}

func (h *WebHandler) IsLoggedIn(r *http.Request) (bool, error) {
	loggedIn := true
	session, err := h.sessionStore.Get(r, "auth")
	if err != nil {
		return false, err
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok {
		loggedIn = false
	}

	if !authenticated {
		loggedIn = false
	}

	return loggedIn, nil
}
