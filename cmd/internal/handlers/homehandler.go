package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode([]byte("hello world")); err != nil {
		return
	}
}
