package flagparser

import (
	"flag"

	"github.com/elkcityhazard/contact-keeper/cmd/internal/config"
)

func ParseFlags(config *config.AppConfig) {

	flag.Parse()

}
