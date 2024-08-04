package handlers

import (
	"net/http"

	"htemplx/app/views/components"
	"htemplx/app/views/pages"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Index("htemplx"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.Login())
}

func Register(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.Register())
}

func ForgetPassword(w http.ResponseWriter, r *http.Request) {
	render(w, r, components.ForgetPassword())
}
