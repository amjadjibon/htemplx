package handlers

import (
	"net/http"

	"htemplx/app/views/pages"
)

func Services(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.Services("htemplx"))
}
