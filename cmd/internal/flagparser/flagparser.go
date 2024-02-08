package flagparser

import (
	"flag"

	"github.com/elkcityhazard/contact-keeper/cmd/internal/config"
)

func ParseFlags(config *config.AppConfig) {

	flag.StringVar(&config.Port, "port", "8080", "Port to listen on")
	flag.StringVar(&config.DSN, "dsn", "", "Database connection string")
	flag.Parse()

}
