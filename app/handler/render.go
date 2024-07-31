package handler

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(
	w http.ResponseWriter,
	r *http.Request,
	comp templ.Component,
) {
	if err := comp.Render(r.Context(), w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
