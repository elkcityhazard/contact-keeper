package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode("stub home route"); err != nil {
		return
	}
}
