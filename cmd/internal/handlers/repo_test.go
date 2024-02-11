package handlers

import (
	"testing"

	"github.com/elkcityhazard/contact-keeper/cmd/internal/config"
)

func Test_GetRepo(t *testing.T) {

	var app config.AppConfig

	NewRepo(&app)

	repo := GetRepo()

	if repo == nil {
		t.Error("Repo is nil")
	}
}

func Test_NewRepo(t *testing.T) {
	var app config.AppConfig
	NewRepo(&app)

	if Repo == nil {
		t.Error("Repo is nil")
	}
}
