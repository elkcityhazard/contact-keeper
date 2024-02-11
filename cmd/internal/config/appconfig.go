package config

import "database/sql"

type AppConfig struct {
	Port    string
	DSN     string
	Version string
	DB      *sql.DB
}
