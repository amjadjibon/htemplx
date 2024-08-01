package handlers

import (
	"net/http"

	"htemplx/app/views/pages"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.Index("htemplx"))
}
