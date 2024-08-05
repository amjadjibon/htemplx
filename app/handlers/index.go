package handlers

import (
	"net/http"

	"htemplx/app/views/components"
	"htemplx/app/views/pages"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Index("htemplx"))
}

func About(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.About("htemplx"))
}

func Contact(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Contact("htemplx"))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, r, pages.NotFound("htemplx"))
}

func Services(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Services("htemplx"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.Login())
}

func Register(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.Register())
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.ForgotPassword())
}

func UnderConstruction(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.UnderConstruction())
}

func TermsAndConditions(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Terms("htemplx"))
}

func Privacy(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Privacy("htemplx"))
}
