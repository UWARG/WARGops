package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/UWARG/WARGops/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

func main() {
	var user goth.User
	port := flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

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
	service := server.NewFinances(db)

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	//TODO: improve cors handling
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
			//TODO: move to env variables
			"1069446618056761375",
			"9-pHUrOkF4pJnPdxFwFQebgtI6mbf5gq",
			"http://localhost:8080/auth/callback?provider=discord",
			"identify", "guilds",
		),
	)

	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		gothic.BeginAuthHandler(w, r)
	})

	r.Get("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err = gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(user.RawData["username"], "has logged in")
		fmt.Println("")
		fmt.Printf("+%v", user)
		fmt.Println("")
		fmt.Println("")

		http.Redirect(w, r, "http://localhost:5173/", http.StatusMovedPermanently)
		// Use the user data however you want.
		// E.g. store in a database, set a session, etc.
	})

	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		respondwithJSON(w, http.StatusOK, user)
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
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
