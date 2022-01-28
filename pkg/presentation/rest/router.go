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
	"drello-api/pkg/util/myerr"
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	fn(w, r)
	err := fn(w, r)
	if err == nil {
		return
	}

	log.Printf("[error]: %v", err)

	httpError, ok := err.(*myerr.HTTPError)
	if !ok {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Unwraped error occured"))
		return
	}

	if 400 <= httpError.Status && httpError.Status < 500 {
		w.WriteHeader(httpError.Status)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(httpError.Detail))
		return
	} else {
		w.WriteHeader(httpError.Status)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Internal server error"))
		return
	}
}
