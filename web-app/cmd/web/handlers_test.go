package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Dummy logger
func newTestApplication() *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

// TODO: had to exclude the home test for now as i have to look into how to include the template files in tests or outsource them
func TestRoutes(t *testing.T) {
	app := newTestApplication()

	tests := []struct {
		name       string
		url        string
		wantStatus int
	}{
		//{"Home", "/", http.StatusOK},
		{"SnippetViewValid", "/snippet/view/1", http.StatusOK},
		{"SnippetViewInvalid String", "/snippet/view/go-tutorial", http.StatusNotFound},
		{"SnippetViewInvalid Negative Int", "/snippet/view/-4", http.StatusNotFound},
		{"SnippetCreate", "/snippet/create", http.StatusOK},
		{"NotFound", "/doesnotexist", http.StatusNotFound},
	}

	router := app.routes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, rr.Code)
			}
		})
	}
}
