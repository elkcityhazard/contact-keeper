package handlers

import (
	"fmt"
	"net/http"
)

func (m *Repository) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "stub get users route")
}
