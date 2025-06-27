// Package handlers provides common utilities for HTTP handlers.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// writeJSONResponse writes a JSON response with the given status code
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

// writeErrorResponse writes an error response
func writeErrorResponse(w http.ResponseWriter, statusCode int, errorType, message string) {
	errorResp := ErrorResponse{
		Error:   errorType,
		Message: message,
		Code:    statusCode,
	}
	writeJSONResponse(w, statusCode, errorResp)
}