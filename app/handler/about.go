package handler

import (
	"htemplx/app/view/pages"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.About("htemplx"))
}
