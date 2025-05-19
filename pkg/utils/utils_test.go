package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prototype01/pkg/utils"
)

func TestJSONResponse(t *testing.T) {
	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the function with test data
	testData := map[string]string{"message": "hello world"}
	utils.JSONResponse(rr, http.StatusOK, testData)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type
	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content-type: got %v want %v", contentType, expected)
	}

	// Check the body
	expected = `{"message":"hello world"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
