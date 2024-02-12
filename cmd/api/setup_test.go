package main

import (
	"testing"

	"github.com/elkcityhazard/contact-keeper/internal/config"
)

var (
	testApp config.AppConfig
)

func TestMain(m *testing.M) {

	testApp = config.AppConfig{
		Port: ":8080",
	}
}
