package handlers

import "net/http"

func (m *Repository) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
