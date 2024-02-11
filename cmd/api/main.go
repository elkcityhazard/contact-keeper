package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/elkcityhazard/contact-keeper/cmd/internal/config"
	"github.com/elkcityhazard/contact-keeper/cmd/internal/flagparser"
	"github.com/elkcityhazard/contact-keeper/cmd/internal/handlers"
	_ "github.com/go-sql-driver/mysql"
)

// flag vars

var (
	DSN     string
	Version string
)

var app config.AppConfig

func main() {

	flagparser.ParseFlags(&app)

	app.DSN = DSN
	app.Version = Version

	db, err := sql.Open("mysql", app.DSN)

	if err != nil {
		log.Fatalln(err)
	}

	app.DB = db

	err = app.DB.Ping()

	if err != nil {
		log.Fatalln(err)
	}

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
