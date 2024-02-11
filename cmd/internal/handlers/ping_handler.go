package handlers

import (
	"encoding/json"
	"net/http"
)

func (m *Repository) PingHandler(w http.ResponseWriter, r *http.Request) {

	type status struct {
		StatusCode int    `json:"status_code"`
		Version    string `json:"version"`
		Message    string `json:"message"`
	}

	var pong status
	pong.StatusCode = 200
	pong.Version = Repo.app.Version
	pong.Message = "pong"

	if err := json.NewEncoder(w).Encode(pong); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
