package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elkcityhazard/contact-keeper/internal/config"
)

func Test_PingHandler(t *testing.T) {

	var app config.AppConfig

	app.Version = "0.0.1"
	app.Port = ":8080"

	NewRepo(&app)

	var pong PingStatus
	pong.StatusCode = 200
	pong.Version = Repo.app.Version
	pong.Message = "pong"

	bytes, err := json.Marshal(pong)

	if err != nil {
		t.Fatalf("could not serialize: %v", err)
	}

	expectedOutput := `{"status_code":200,"version":"0.0.1","message":"pong"}`

	if string(bytes) != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bytes), expectedOutput)
	}

	// create the request

	req, err := http.NewRequest("GET", "/ping", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// create a response recorder

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

	t.Log(rr)

	responseBody := PingStatus{}

	err = json.Unmarshal([]byte(rr.Body.Bytes()), &responseBody)

	if err != nil {
		t.Fatal(err)
	}

	// check the response body

	if responseBody != pong {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, pong)
	}

}
