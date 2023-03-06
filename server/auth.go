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

func (s Server) SignInAsGuest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signing in as guest")
	user := goth.User{
		UserID:    "guest",
		Name:      "Guest",
		AvatarURL: "https://cdn.discordapp.com/embed/avatars/0.png",
		RawData: map[string]interface{}{
			"roles": []string{"guest"},
		},
	}

	store, err := s.users.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.Save(r, w)
	store.Values["data"] = user
	store.Save(r, w)
	json.NewEncoder(w).Encode(user)
}

func (s Server) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging out")
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
	user, ok := r.Context().Value(UserKey).(goth.User)
	if !ok {
		http.Error(w, "could not find user", http.StatusBadRequest)
	}

	member, err := s.bot.GuildMember(Config.DiscordGuildID, user.UserID)
	if err != nil {
		if user.UserID == "guest" {
			enc := json.NewEncoder(w)
			enc.SetIndent("", "\t")
			enc.Encode(user)
			return
		}
		http.Error(w, "could not find roles: "+err.Error(), http.StatusBadRequest)
		return
	}

	user.AccessToken = ""
	user.RefreshToken = ""
	user.Name = member.User.String()
	user.RawData["roles"] = member.Roles
	user.NickName = member.Nick

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(user)
}
