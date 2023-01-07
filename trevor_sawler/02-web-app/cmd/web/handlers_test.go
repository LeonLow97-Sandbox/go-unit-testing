package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name string
		url string
		expectedStatusCode int
	} {
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	// create a test server (start a server as part of testing)
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// inside cmd/web, go up 2 levels to reach the templates folder when running tests
	pathToTemplates = "./../../templates/"

	// range through test data
	// ts.URL is the url that the test web server has and append to it the specified endpoint
	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}