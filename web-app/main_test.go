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

func TestSnippetView(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/snippet/view", nil)
	rr := httptest.NewRecorder()

	snippetView(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected OK, got %d", rr.Code)
	}

	expected := "Display specific snippet"
	if rr.Body.String() != expected {
		t.Errorf("expected response %q, got %q", expected, rr.Body.String())
	}
}

func TestSnippetCreate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/snippet/create", nil)
	rr := httptest.NewRecorder()

	snippetCreate(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected OK, got %d", rr.Code)
	}

	expected := "Display form to create new snippet"
	if rr.Body.String() != expected {
		t.Errorf("expected response %q, got %q", expected, rr.Body.String())
	}
}
