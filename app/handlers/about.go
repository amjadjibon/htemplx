package handlers

import (
	"net/http"

	"htemplx/app/views/pages"
)

func About(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.About("htemplx"))
}
