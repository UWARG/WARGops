package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/UWARG/WARGops/bot"
	"github.com/UWARG/WARGops/bot/commands"
	"github.com/UWARG/WARGops/server"
	"github.com/bwmarrin/discordgo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

type Config struct {
	// Discord
	DiscordBotToken     string `json:"discord_bot_token"`
	DiscordClientID     string `json:"discord_client_id"`
	DiscordClientSecret string `json:"discord_client_secret"`
	DiscordRedirectURL  string `json:"discord_redirect_url"`
}

var config Config

func main() {
	ReadConfig()
	var user goth.User

	port := flag.Int("port", 8080, "Port for test HTTP server")
	flag.StringVar(&commands.LeadRoleID, "leads", "820466330456162354", "team lead role")
	flag.StringVar(&commands.WorkspaceID, "workspace", "71624493506711", "asana workspace ID")
	botToken := flag.String("discord", "", "discord bot token")
	asanaToken := flag.String("asana", "", "discord personal access token")
	flag.Parse()

	if *botToken != "" {
		go func() {
			bot.Start(*botToken, *asanaToken)
		}()
	}

	swagger, err := server.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	db, err := server.NewDB("db.sqlite")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open database: %v", err)
		os.Exit(1)
	}

	discordBot, err := discordgo.New("Bot " + config.DiscordBotToken)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to Bot: %v", err)
		os.Exit(1)
	}

	service := server.NewFinances(db, discordBot)

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// TODO: improve cors handling
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// AUTHENTICATION
	goth.UseProviders(
		discord.New(
			// TODO: move to env variables
			config.DiscordClientID,
			config.DiscordClientSecret,
			config.DiscordRedirectURL,
			"identify",
		),
	)

	key := "test"        // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store // Set the store as the session store

	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		gothic.BeginAuthHandler(w, r)
	})

	r.Get("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err = gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "http://localhost:5173/", http.StatusMovedPermanently)
	})

	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		body := getFromDiscord("https://discord.com/api/users/@me", user.AccessToken, w)
		var data map[string]interface{}
		json.NewDecoder(body.Body).Decode(&data)
		respondwithJSON(w, http.StatusOK, data)
	})

	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging out")
		user = goth.User{}
		gothic.Logout(w, r)
	})

	server.Handler(service, server.WithRouter(r))

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	fmt.Printf("Starting server on port http://localhost:%d/\n", *port)

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getFromDiscord(endpoint string, token string, w http.ResponseWriter) *http.Response {
	//Create the request for the user endpoint
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	//Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return resp
}

func ReadConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

}
