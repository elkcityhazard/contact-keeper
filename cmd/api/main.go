package main

import (
	"context"
	"database/sql"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/contact-keeper/internal/config"
	"github.com/elkcityhazard/contact-keeper/internal/flagparser"
	"github.com/elkcityhazard/contact-keeper/internal/handlers"
	"github.com/elkcityhazard/contact-keeper/internal/mailer"
	"github.com/elkcityhazard/contact-keeper/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

// flag vars

var (
	DSN          string
	Version      string
	RediPassword string
)

var app config.AppConfig

func main() {

	gob.Register(models.User{})
	gob.Register(handlers.PingStatus{})

	flagparser.ParseFlags(&app)

	// connect to redis

	redisClient := NewRedisPool()
	app.RedisPool = redisClient

	sessionManager := scs.New()
	sessionManager.Store = redisstore.New(redisClient)

	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = true // Set to true in production

	app.DSN = DSN
	app.Version = Version
	app.Context = context.Background()
	app.WG = &sync.WaitGroup{}
	app.Mutex = &sync.Mutex{}
	app.SessionManager = sessionManager

	app.ErrorChan = make(chan error)
	app.ErrorDoneChan = make(chan bool)

	db, err := sql.Open("mysql", app.DSN)

	if err != nil {
		log.Fatalln(err)
	}

	mailer.NewMailerConfig(&app)

	app.DB = db
	app.DBTimeout = 10 * time.Second

	err = app.DB.Ping()

	if err != nil {
		log.Fatalln(err)
	}

	// listen for shutdown

	go listenForShutdown()
	go app.ListenForErrors()

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

func NewRedisPool() *redis.Pool {

	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			// redis password setup

			if RediPassword != "" {
				if _, err := c.Do("AUTH", RediPassword); err != nil {
					return nil, err
				}
			}
			return c, err
		},
	}
}

func listenForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.Shutdown()
	os.Exit(0)
}
