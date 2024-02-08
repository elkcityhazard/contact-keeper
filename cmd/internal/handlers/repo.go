package handlers

import "github.com/elkcityhazard/contact-keeper/cmd/internal/config"

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repo with the given AppConfig.
//
// It takes a pointer to an AppConfig as a parameter and returns a pointer to a Repo.
func NewRepo(a *config.AppConfig) {
	Repo = &Repository{
		app: a,
	}
}

func GetRepo() *Repository {
	return Repo
}
