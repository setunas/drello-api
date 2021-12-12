package rest

import (
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")

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
	router.Handle("/workspaces/{id:[0-9]+}", handler(workspaceHandler))
	router.Handle("/workspaces", handler(workspacesHandler))
	router.Handle("/boards/{id:[0-9]+}", handler(boardHandler))
	router.Handle("/columns/{id:[0-9]+}", handler(columnHandler))
	router.Handle("/columns", handler(columnsHandler))
	router.Handle("/cards/{id:[0-9]+}", handler(cardHandler))
	router.Handle("/cards", handler(cardsHandler))
}
