package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func SetupGoth(store sessions.Store) {
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("GOOGLE_CALLBACK"),
		),
	)

	gothic.Store = store
}

func GothicLogin(res http.ResponseWriter, req *http.Request) (goth.User, error) {
	return gothic.CompleteUserAuth(res, req)
}

func GothicBeginAuthHandler(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}
func GothicLogout(res http.ResponseWriter, req *http.Request) error {
	return gothic.Logout(res, req)
}
