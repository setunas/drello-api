package rest

import (
	"drello-api/pkg/presentation/controller"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/workspaces", controller.Workspaces)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
