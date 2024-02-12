package flagparser

import (
	"flag"

	"github.com/elkcityhazard/contact-keeper/internal/config"
)

func ParseFlags(config *config.AppConfig) {

	flag.Parse()

}
