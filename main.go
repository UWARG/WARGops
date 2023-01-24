package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/UWARG/WARGops/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

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

	server.Handler(service, server.WithRouter(r))

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	fmt.Printf("Starting server on port http://localhost:%d/", *port)

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
