package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"            // lightweight HTTP router
	"github.com/go-chi/chi/v5/middleware"  // built-in middlewares (logging, recovery, etc.)
)

// application holds all app-level dependencies
type application struct {
	config config
}

// config stores the server and database settings
type config struct {
	addr string   // server address, e.g. ":8080"
	db   dbconfig // database settings
}

// dbconfig holds the database connection string
type dbconfig struct {
	dsn string // data source name used to connect to the DB
}

// run starts the HTTP server and registers routes
func (app *application) run() error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)    // logs every incoming request
	r.Use(middleware.Recoverer) // catches panics and returns 500 instead of crashing

	// simple health-check endpoint — returns "OK" if server is alive
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Printf("Starting server on %s", app.config.addr)
	return http.ListenAndServe(app.config.addr, r) // blocks and serves until error
}
