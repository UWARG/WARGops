package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (s Server) Authenticate(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (s Server) Callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store, err := s.users.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.Save(r, w)
	store.Values["data"] = user
	store.Save(r, w)

	http.Redirect(w, r, Config.FrontendURI, http.StatusTemporaryRedirect)
	json.NewEncoder(w).Encode(user)
}

func (s Server) Logout(w http.ResponseWriter, r *http.Request) {
	// Logout from the provider
	gothic.Logout(w, r)

	// Clear the sotre
	store, err := s.users.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Values["data"] = nil
	store.Save(r, w)
	// http.Redirect(w, r, Config.FrontendURI, http.StatusTemporaryRedirect)
}

func (s Server) Info(w http.ResponseWriter, r *http.Request) {
	store, err := s.users.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(s.users)

	raw, ok := store.Values["data"]
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(raw)

	user, ok := raw.(goth.User)
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}
