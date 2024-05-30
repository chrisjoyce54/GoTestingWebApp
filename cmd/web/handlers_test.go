package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"home", "/fish", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	pathToTemplates = "./../../templates/"
	for _, test := range tests {
		resp, err := testServer.Client().Get(testServer.URL + test.url)

		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != test.expectedStatusCode {
			t.Errorf("For %s: expected status %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
		}
	}
}
