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

type handler func(http.ResponseWriter, *http.Request)

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHTTPRequest(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	fn(w, r)
}
