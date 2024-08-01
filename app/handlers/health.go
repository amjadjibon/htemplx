package handlers

import "net/http"

func Healthz(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("{}"))
}
