package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routes(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		handlerFunc  http.HandlerFunc
		expectedCode int
		expectedBody string
	}{
		{
			name:   "home page",
			method: "GET",
			path:   "/",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				fmt.Println(r)
				w.Write([]byte("Hello, From Home Page"))
			},
			expectedCode: 200,
			expectedBody: "Hello, From Home Page",
		},
		{
			name:   "test api route",
			method: "GET",
			path:   "/ping",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			},
			expectedCode: 200,
			expectedBody: "pong",
		},
		{
			name:   "test api post user route",
			method: "GET",
			path:   "/api/users",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("post to user endpoint"))
			},
			expectedCode: 200,
			expectedBody: "post to user endpoint",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := chi.NewRouter()

			router.Method(tc.method, tc.path, tc.handlerFunc)

			server := httptest.NewServer(router)

			defer server.Close()

			resp, err := http.Get(server.URL + tc.path)
			if err != nil {
				t.Fatalf("Error while making http request: %v", err)
			}

			// TODO: add more test cases especially POST PUT DELETE

			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Fatalf("Expected status %d, got %d", tc.expectedCode, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error while reading response: %v", err)
			}

			if strings.Trim(string(body), "\n") != tc.expectedBody {
				t.Fatalf("Expected body %s, got %s", tc.expectedBody, string(body))
			}
		})
	}
}
