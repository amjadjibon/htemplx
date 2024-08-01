package handlers

import (
	"net/http"

	"htemplx/app/views/pages"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.Contact("htemplx"))
}
