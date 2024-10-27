package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the root handler
func TestGetDetailHandler_MethodNotAllowed(t *testing.T) {
    // هنا نختبر إذا كانت الطريقة غير GET
    req, err := http.NewRequest("get", "/artist/1", nil) // نستخدم POST بدلاً من GET
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(getHandler)

    handler.ServeHTTP(rr, req)

    // نتحقق أن الاستجابة هي 405 Method Not Allowed
    if status := rr.Code; status != http.StatusMethodNotAllowed {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusMethodNotAllowed)
    }

    expected := "Method Not Allowed"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
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
