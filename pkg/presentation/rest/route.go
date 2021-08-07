package rest

import (
	"drello-api/pkg/constants"
	"fmt"
	"log"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHTTPRequest(r)
	fn(w, r)
}

func HandleRequests() {
	fmt.Println("Listening on http://127.0.0.1:8080")

	http.Handle(constants.Workspaces, handler(workspaceHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
