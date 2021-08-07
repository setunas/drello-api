package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type clientErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func newClientErrorResponse(status int, message string) *clientErrorResponse {
	return &clientErrorResponse{Status: status, Message: message}
}

func handleClientError(w http.ResponseWriter, err error, status int, message string) {
	log.Printf("REST Client Error: %v", err)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(newClientErrorResponse(status, message))
}

func handleServerError(w http.ResponseWriter, err error) {
	log.Printf("REST Server Error: %v", err)
	w.WriteHeader(500)
}
