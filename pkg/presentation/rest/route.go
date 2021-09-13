package rest

import (
	"drello-api/pkg/constants"
	"drello-api/pkg/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

type handler func(http.ResponseWriter, *http.Request)

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHTTPRequest(r)
	fn(w, r)
}

func HandleRequests() {
	port := fmt.Sprintf(":%s", utils.MustGetenv("PORT"))
	fmt.Println("Listening on PORT:", port)

	router = mux.NewRouter()
	setHandlers()
	log.Fatal(http.ListenAndServe(port, router))
}

func setHandlers() {
	router.Handle(constants.Workspaces+"/{id:[0-9]+}", handler(workspaceHandler))
	router.Handle(constants.Workspaces, handler(workspacesHandler))
}
