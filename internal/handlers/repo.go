package handlers

import (
	"encoding/gob"

	"github.com/elkcityhazard/contact-keeper/internal/config"
	"github.com/elkcityhazard/contact-keeper/internal/models"
)

type status struct {
	StatusCode int
	Message    string
}

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repo with the given AppConfig.
//
// It takes a pointer to an AppConfig as a parameter and returns a pointer to a Repo.
func NewRepo(a *config.AppConfig) {
	gob.Register(status{})
	Repo = &Repository{
		app: a,
	}
}

func GetRepo() *Repository {
	return Repo
}

func (m *Repository) GetUsers() ([]*models.User, error) {
	return nil, nil
}
