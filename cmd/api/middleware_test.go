package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_middleware(t *testing.T) {

	// create the tests

	pathTests := []struct {
		name               string
		path               string
		expectedRedirect   string
		expectedStatusCode int
	}{
		{
			name:               "home page",
			path:               "/",
			expectedRedirect:   "/",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "test ping route redirect",
			path:               "/ping/",
			expectedRedirect:   "/ping",
			expectedStatusCode: http.StatusMovedPermanently,
		},
		{
			name:               "test ping ok status code",
			path:               "/ping",
			expectedRedirect:   "/ping",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "test random slug",
			path:               "/api/v1/random-post/this-is-a-test/",
			expectedRedirect:   "/api/v1/random-post/this-is-a-test",
			expectedStatusCode: http.StatusMovedPermanently,
		},

		{
			name:               "test random slug",
			path:               "/api/v1/random-post/this-is-a-test///",
			expectedRedirect:   "/api/v1/random-post/this-is-a-test",
			expectedStatusCode: http.StatusMovedPermanently,
		},
	}

	// loop through the tests

	for _, tt := range pathTests {
		// run the test
		t.Run(tt.name, func(t *testing.T) {

			r := chi.NewRouter()

			r.Use(StripTrailingSlash)

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, From Home Page"))
			})

			r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			})

			r.Get("/api/v1/random-post/{slug}", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			})

			// start up a server with chi router
			ts := httptest.NewServer(r)
			defer ts.Close()

			nextURL := ts.URL + tt.path

			for i := 0; i < 10; i++ {

				client := &http.Client{

					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						fmt.Println(req.URL.String())
						return http.ErrUseLastResponse
					},
				}

				req, err := http.NewRequest("GET", nextURL, nil)
				if err != nil {
					t.Fatal(err)
				}

				// make the request
				resp, err := client.Do(req)
				if err != nil {
					t.Fatal(err)
				}
				defer resp.Body.Close()

				if resp.StatusCode == 200 {
					if resp.Request.URL.Path != tt.expectedRedirect {
						t.Errorf("expected redirect %s, got %s", tt.path, resp.Request.URL.Path)
					}
					break
				} else {

					log.Println("statusCode: ", resp.StatusCode)
					// check the status code
					if resp.StatusCode != tt.expectedStatusCode {
						t.Errorf("expected status code %d, got %d", tt.expectedStatusCode, resp.StatusCode)
					}

					// check the redirect
					if resp.Request.URL.Path != tt.path {
						t.Errorf("expected path %s, got %s", tt.path, resp.Request.URL.Path)
					}

					nextURL = resp.Request.URL.String()
					i++
				}

			}

		})
	}

}
