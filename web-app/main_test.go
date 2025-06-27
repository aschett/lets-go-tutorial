package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	id := 20
	url := fmt.Sprintf("/snippet/view/%d", id)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.SetPathValue("id", strconv.Itoa(id))
	rr := httptest.NewRecorder()

	snippetView(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected OK, got %d", rr.Code)
	}

	expected := fmt.Sprintf("Display a specific snipped with ID %d", id)
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
