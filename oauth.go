package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/browser"
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
	randomState = "state"
)

//initOauth2 ...
func initOauth2() {
	http.HandleFunc("/oauth", handleAuthentification)
	url := googleOauthConfig.AuthCodeURL(randomState, oauth2.AccessTypeOffline)
	browser.OpenURL(url)
	http.ListenAndServe(":8080", nil)
}

func handleAuthentification(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("State is not valid!")
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Println("Could not fetch the token: %s\n", err.Error())
		return
	}
	os.Setenv("TOKEN_GTASK", token.AccessToken)
	if os.Getenv("TOKEN_GTASK") == "" {
		fmt.Println("Token has not been fetched successfully, please try again by running the command LOGIN")
	} else {
		fmt.Println("Token has been fetched successfully: " + token.AccessToken)
	}
}
