package main

import (
	"context"
	"fmt"
	_ "golang.org/x/net/context"
	"net/http"
	"os"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8081/oauth",
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

	err := browser.OpenURL(url)
	if err != nil {
		fmt.Println("Error opening the browser", err)
		os.Exit(1)
	}

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting gtask receiver endpoint (port:8081)", err)
		os.Exit(1)
	}
}

func handleAuthentification(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("State is not valid!")
		return
	}
	token, err := googleOauthConfig.Exchange(context.TODO(), r.FormValue("code"))
	if err != nil {
		fmt.Println("Could not fetch the token. Make sure you have access to the internet", err.Error())
		return
	}
	err = os.Setenv("TOKEN_GTASK", token.AccessToken)
	if err != nil {
		fmt.Println("Error creating the environment variable", err)
		os.Exit(1)
	}
	if os.Getenv("TOKEN_GTASK") == "" {
		fmt.Println("Token has not been fetched successfully, please try again by running the command LOGIN")
	} else {
		fmt.Println("Token has been fetched successfully: " + token.AccessToken)
	}
}
