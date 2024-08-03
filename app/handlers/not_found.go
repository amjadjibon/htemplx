package handlers

import (
	"net/http"

	"htemplx/app/views/pages"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.NotFound("htemplx"))
}
