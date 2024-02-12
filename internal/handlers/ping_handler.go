package handlers

import (
	"encoding/json"
	"net/http"
)

type PingStatus struct {
	StatusCode int    `json:"status_code"`
	Version    string `json:"version"`
	Message    string `json:"message"`
}

func (m *Repository) PingHandler(w http.ResponseWriter, r *http.Request) {

	msg := m.app.SessionManager.GetString(r.Context(), "ping")
	if msg != "" {
		msg = "no session currently, retry this request"
	}

	var pong PingStatus
	pong.StatusCode = 200
	pong.Version = Repo.app.Version
	pong.Message = msg

	m.app.SessionManager.Put(r.Context(), "ping", pong)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(pong); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
