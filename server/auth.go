package server

import (
	"encoding/json"
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

	store.Values["data"] = user
	store.Save(r, w)
	http.Redirect(w, r, Config.FrontendURI, http.StatusTemporaryRedirect)
	json.NewEncoder(w).Encode(user)
}

func (s Server) Logout(w http.ResponseWriter, r *http.Request) {
}

func (s Server) Info(w http.ResponseWriter, r *http.Request) {
	store, err := s.users.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	raw, ok := store.Values["data"]
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, ok := raw.(goth.User)
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}
