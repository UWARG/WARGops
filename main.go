package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

func init() {
	buf, err := os.ReadFile("config.json")
	if err != nil {
		panic(fmt.Errorf("could not find config.json file: %v", buf))
	}

	if err := json.Unmarshal(buf, &server.Config); err != nil {
		panic(err)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	config := server.Config
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
		return err
	}
	swagger.Servers = nil

	db, err := server.NewDB("db.sqlite")
	if err != nil {
		return err
	}

	discordBot, err := discordgo.New("Bot " + config.DiscordBotToken)
	if err != nil {
		return err
	}

	service := server.NewFinances(db, discordBot)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://ops.uwarg.com", "http://localhost:5173", "http://ops.uwarg.com"},
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

	store := sessions.NewCookieStore([]byte(config.Secret))
	store.MaxAge(3600) // 1 hour
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false
	gothic.Store = store

	r.Route("/api", func(r chi.Router) {
		server.Handler(service, server.WithRouter(r))
		r.Get("/auth", service.Authenticate)
		r.Get("/auth/callback", service.Callback)
		r.Get("/info", service.Info)
		r.Get("/logout", service.Logout)
	})

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("/home/h/code/WARGops/front/dist/assets"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./front/dist/index.html")
	})

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	fmt.Printf("Starting server on port http://localhost:%d/\n", *port)
	return s.ListenAndServe()
}
