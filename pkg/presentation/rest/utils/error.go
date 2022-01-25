package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(message string) *errorResponse {
	return &errorResponse{Message: message}
}

func HandleClientError(w http.ResponseWriter, err error, status int, message string) {
	log.Printf("REST Client Error: %v", err)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(newErrorResponse(message))
}

func HandleServerError(w http.ResponseWriter, err error) {
	log.Printf("REST Server Error: %v", err)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(newErrorResponse("An internal server error occurred"))
}
