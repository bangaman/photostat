package callback

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"

	"app"
	"auth"
	"fmt"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query()["state"][0] != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	fmt.Println(profile)
	get_sub := profile["sub"].(string)

	//profile is a map with different key => values
	//it can be user details directlty from 

	if get_sub[:13] == "google-oauth2" {
		fmt.Println(profile["given_name"])
		session.Values["username"] = profile["given_name"]
		session.Values["profilepic"] = profile["picture"]
		session.Values["email"]  = profile["nickname"]
		session.Values["usersid"] = profile["sub"]
	}else{

		session.Values["usersid"] = profile["sub"]
		session.Values["email"] = profile["name"]
		session.Values["profilepic"] = profile["picture"]
		session.Values["username"] = profile["nickname"]
	}

	fmt.Println(profile)
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to admin for verification
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
