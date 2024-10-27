package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the root handler
func TestGetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the response code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check for expected content in the response body
	expected := "<h1>List of Artists</h1>"
	if rr.Body.String()[:len(expected)] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Test the detail handler with a valid artist ID
func TestGetDetailHandler_ValidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/detail?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getDetail)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the response code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check for expected content in the response body
	expected := "<img src=\""
	if rr.Body.String()[:len(expected)] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Test the detail handler with an invalid artist ID
func TestGetDetailHandler_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/detail?id=", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getDetail)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the response code for bad request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

// Test the error handler
func TestErrorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/invalid-path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the response code for not found
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
