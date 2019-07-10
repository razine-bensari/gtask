package main

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/oauth",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/tasks"},
		Endpoint:     google.Endpoint,
	}
	state = "state"
	//AccessTypeOffline ...
	AccessTypeOffline = SetAuthURLParam("access_type", "offline")
)

//initOauth2 ...
func initOauth2() {
	http.HandleFunc("/oauth", handleAuthentification())
	http.ListenAndServe(":8080", nil)
}

func handleAuthentification(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(state, AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//add handle the pop page for oath
