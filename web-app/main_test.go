package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	home(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected OK, got %d", rr.Code)
	}

	expected := "Hello from andibox"
	if rr.Body.String() != expected {
		t.Errorf("expected response %q, got %q", expected, rr.Body.String())
	}
}
