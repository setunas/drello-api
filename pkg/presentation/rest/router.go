package rest

import (
	"drello-api/pkg/presentation/rest/boardHandler"
	"drello-api/pkg/presentation/rest/cardHandler"
	"drello-api/pkg/presentation/rest/cardPositionsHandler"
	"drello-api/pkg/presentation/rest/cardsHandler"
	"drello-api/pkg/presentation/rest/columnHandler"
	"drello-api/pkg/presentation/rest/columnsHandler"
	"drello-api/pkg/presentation/rest/meHandler"
	"drello-api/pkg/presentation/rest/signupHandler"
	"drello-api/pkg/util"
	"drello-api/pkg/util/apperr"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func HandleRequests() {
	port := fmt.Sprintf(":%s", util.MustGetenv("PORT"))
	fmt.Println("Listening on PORT:", port)

	router = mux.NewRouter()
	setHandlers()
	log.Fatal(http.ListenAndServe(port, router))
}

func setHandlers() {
	router.Handle("/me", handler(meHandler.MeHandler))
	router.Handle("/signup", handler(signupHandler.SignupHandler))
	router.Handle("/boards/{id:[0-9]+}", handler(boardHandler.BoardHandler))
	router.Handle("/columns/{id:[0-9]+}", handler(columnHandler.ColumnHandler))
	router.Handle("/columns", handler(columnsHandler.ColumnsHandler))
	router.Handle("/cards/positions", handler(cardPositionsHandler.CardPositionsHandler))
	router.Handle("/cards/{id:[0-9]+}", handler(cardHandler.CardHandler))
	router.Handle("/cards", handler(cardsHandler.CardsHandler))
}

type handler func(http.ResponseWriter, *http.Request) error

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHTTPRequest(r)
	setHeaders(w)

	fn(w, r)
	err := fn(w, r)
	if err == nil {
		return
	}

	handleError(w, err)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func handleError(w http.ResponseWriter, err error) {
	type errRes struct {
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")

	httpError, ok := err.(*apperr.HTTPError)
	if !ok {
		log.Printf("[error]: Unknown error type: %v", err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errRes{Message: "Unknown error occured"})
		return
	}

	log.Printf("[error]: %v: occurred at %s", httpError, httpError.OccurredAt())
	w.WriteHeader(httpError.Status())
	if httpError.IsClientError() {
		json.NewEncoder(w).Encode(errRes{Message: httpError.Error()})
	} else {
		json.NewEncoder(w).Encode(errRes{Message: "Internal server error"})
	}
}
