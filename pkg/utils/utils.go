// Package utils provides common utility functions used throughout the application
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TimeFormat is the standard time format used in the application
const TimeFormat = time.RFC3339

// JSONResponse sends a JSON response with the given status code and data
func JSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Error marshaling JSON response"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// SendErrorResponse is a helper function to send error responses
func SendErrorResponse(w http.ResponseWriter, code int, errMessage string, message string) {
	JSONResponse(w, code, ErrorResponse{
		Error:   errMessage,
		Message: message,
		Code:    code,
	})
}

// FormatTimestamp formats a time.Time into a string using standard format
func FormatTimestamp(t time.Time) string {
	return t.Format(TimeFormat)
}

// ParseTimestamp parses a timestamp string into a time.Time using standard format
func ParseTimestamp(s string) (time.Time, error) {
	return time.Parse(TimeFormat, s)
}
