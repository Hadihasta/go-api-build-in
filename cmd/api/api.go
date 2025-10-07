package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hadihasta/go-api-build-in/store"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {

// func (app *application) mount() *http.ServeMux {
	// using mux from standard library
	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	// return mux

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Get("/health", app.healthCheckHandler)

	r.Route("/v1", func(r chi.Router){
		r.Get("/health", app.healthCheckHandler)
	})
	return r
	
}

// func (app *application) run(mux *http.ServeMux) error {
func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}
	log.Printf("starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
