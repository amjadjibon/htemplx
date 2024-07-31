package handler

import (
	"htemplx/app/view/pages"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, r, pages.Index("Home"))
}
