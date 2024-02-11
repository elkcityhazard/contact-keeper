package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elkcityhazard/contact-keeper/cmd/internal/config"
)

func Test_PingHandler(t *testing.T) {

	var app config.AppConfig

	app.Version = "0.0.1"
	app.Port = ":8080"

	NewRepo(&app)

	// create the request

	req, err := http.NewRequest("GET", "/ping", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// create a response recorder

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type status struct {
			StatusCode int    `json:"status_code"`
			Version    string `json:"version"`
			Message    string `json:"message"`
		}

		t.Log(Repo)

		var pong status
		pong.StatusCode = 200
		pong.Version = Repo.app.Version
		pong.Message = "pong"

		if err := json.NewEncoder(w).Encode(pong); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// server the request

	handler.ServeHTTP(rr, req)

	// check the status code

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	// check the response body

	type status struct {
		StatusCode int    `json:"status_code"`
		Version    string `json:"version"`
		Message    string `json:"message"`
	}

	expected := &status{
		StatusCode: 200,
		Version:    "0.0.1",
		Message:    "pong",
	}

	var actual status
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatal(err)
	}

	log.Printf("%v+\n", actual)

	if actual != *expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}

}
