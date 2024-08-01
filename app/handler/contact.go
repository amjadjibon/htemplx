package handler

import (
	"net/http"

	"htemplx/app/view/pages"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.Contact("htemplx"))
}
