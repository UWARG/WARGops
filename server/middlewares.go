package server

import (
	"context"
	"net/http"

	"github.com/markbates/goth"
)

type ContextKey string

var UserKey ContextKey = "user"

func (s Server) HasAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store, err := s.users.Get(r, "auth")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		raw, ok := store.Values["data"]
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, ok := raw.(goth.User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
