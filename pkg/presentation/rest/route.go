package rest

import (
	"drello-api/pkg/constants"
	"fmt"
	"log"
	"net/http"
)

func HandleRequests() {
	fmt.Println("Listening on http://127.0.0.1:8080")

	http.HandleFunc(constants.Workspaces, Workspaces)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
