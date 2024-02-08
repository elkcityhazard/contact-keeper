package main

import (
	"log"
	"net/http"
	"time"

	"github.com/elkcityhazard/contact-keeper/cmd/internal/config"
	"github.com/elkcityhazard/contact-keeper/cmd/internal/flagparser"
	"github.com/elkcityhazard/contact-keeper/cmd/internal/handlers"
)

var app config.AppConfig

func main() {

	flagparser.ParseFlags(&app)

	// create a new repo with the AppConfig

	handlers.NewRepo(&app)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      routes(),
		IdleTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Println("Listening on port 8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server")
		log.Fatal(err)
	}
}
