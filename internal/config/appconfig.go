package config

import (
	"context"
	"database/sql"
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
}
