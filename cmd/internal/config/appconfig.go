package config

import "database/sql"

type AppConfig struct {
	Port string
	DSN  string
	DB   *sql.DB
}
