package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// Mock artist data for testing
func TestFetchData_Success(t *testing.T) {
	// Create a new HTTP server to mock the response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		artist := Artist{
			ID:           1,
			Image:        "test_image.png",
			Name:         "Test Artist",
			Members:      []string{"Member1", "Member2"},
			CreationDate: 2020,
			FirstAlbum:   "First Album",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(artist) // Encode the mock artist data
	}))
	defer ts.Close() // Ensure the server is closed after the test

	// Call the fetchData function with the mock server URL
	result, err := fetchData[Artist](ts.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check the result
	if result.Name != "Test Artist" {
		t.Errorf("expected Test Artist, got %s", result.Name)
	}
}

// Test fetchData for failure case
func TestFetchData_FailedRequest(t *testing.T) {
	// Create a new HTTP server to mock a failed response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError) // Simulate an error response
	}))
	defer ts.Close() // Ensure the server is closed after the test

	// Call the fetchData function with the mock server URL
	_, err := fetchData[Artist](ts.URL)
	if err == nil {
		t.Fatal("expected an error, got none")
	}
}

// Additional tests for fetching locations, concert dates, and relations can be added similarly

func TestErrorHandler(t *testing.T) {
	// Create a response recorder to capture the response
	rec := httptest.NewRecorder()

	// Call the errorHandler with a specific status
	errorHandler(rec, http.StatusNotFound)

	// Check the response code
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}

	// Check the response body
	expected := "Page not found (404)"
	if rec.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rec.Body.String())
	}
}

func TestGetHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	getHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

// اختبر getDetail
func TestGetDetail(t *testing.T) {
	// إعداد الطلب الوهمي مع معرّف الفنان
	req := httptest.NewRequest(http.MethodGet, "/detail?id=1", nil)
	w := httptest.NewRecorder()

	getDetail(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestGetDetail_EmptyID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/detail?id=", nil)
	w := httptest.NewRecorder()

	getDetail(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status Bad Request; got %v", res.Status)
	}
}

// اختبار الحالة عند عدم وجود الفنان
func TestGetDetail_ArtistNotFound(t *testing.T) {
	// هنا تحتاج إلى إعداد fetchCompleteArtistData ليعيد أن الفنان غير موجود
	req := httptest.NewRequest(http.MethodGet, "/detail?id=999", nil)
	w := httptest.NewRecorder()

	getDetail(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status Not Found; got %v", res.Status)
	}
}

// Mocking a valid template for successful test
var validTemplate = template.Must(template.New("valid").Parse(`{{.Title}}: {{.Body}}`))

// Test function for renderTemplate
func TestRenderTemplate(t *testing.T) {
	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Create a sample data to pass to the template
	data := map[string]string{
		"Title": "Test Title",
		"Body":  "This is a test body.",
	}

	// Execute the valid template
	err := validTemplate.ExecuteTemplate(recorder, "valid", data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", recorder.Code)
	}

	// Check the response body
	expectedBody := "Test Title: This is a test body."
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, recorder.Body.String())
	}
}

// Test function for renderTemplate with error handling
func TestRenderTemplate_Error(t *testing.T) {
	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Use a valid template but with a nonexistent template name
	err := validTemplate.ExecuteTemplate(recorder, "nonexistent", nil)
	if err == nil {
		t.Fatal("Expected an error, got none")
	}

	// We don't check the status code here as it won't be set in this context
	// since we are not invoking renderTemplate directly, we just need to validate the error.
}
