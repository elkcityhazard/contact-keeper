package config

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

type AppConfig struct {
	Port           string
	DSN            string
	Version        string
	DB             *sql.DB
	DBTimeout      time.Duration
	Context        context.Context
	WG             *sync.WaitGroup
	Mutex          *sync.Mutex
	RedisPool      *redis.Pool
	SessionManager *scs.SessionManager
	ErrorChan      chan error
	ErrorDoneChan  chan bool
	MailerChan     chan Message
}

func (app *AppConfig) Shutdown() {

	log.Println("Shutting down server...")
	app.WG.Wait()

	// send close to any done chans

	app.ErrorDoneChan <- true

	// close chans

	close(app.ErrorChan)
	close(app.ErrorDoneChan)

	log.Println("closing channels and shutting down app...")

}

func (app *AppConfig) ListenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			if err != nil {
				log.Println(err)
			}
		case <-app.ErrorDoneChan:
			return
		}
	}
}
