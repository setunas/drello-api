package rest

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/handlers/boardHandler"
	"drello-api/pkg/presentation/rest/handlers/cardHandler"
	"drello-api/pkg/presentation/rest/handlers/cardPositionsHandler"
	"drello-api/pkg/presentation/rest/handlers/cardsHandler"
	"drello-api/pkg/presentation/rest/handlers/columnHandler"
	"drello-api/pkg/presentation/rest/handlers/columnsHandler"
	"drello-api/pkg/presentation/rest/handlers/meHandler"
	"drello-api/pkg/presentation/rest/handlers/signupHandler"
	"drello-api/pkg/util"
	"drello-api/pkg/util/log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func HandleRequests() {
	port := fmt.Sprintf(":%s", util.MustGetenv("PORT"))
	log.Info("Listening on PORT:", port).Write()

	router = mux.NewRouter()
	setHandlers()
	log.Fatal(http.ListenAndServe(port, router))
}

func setHandlers() {
	router.Handle("/me", handler.Handler(meHandler.MeHandler))
	router.Handle("/signup", handler.Handler(signupHandler.SignupHandler))
	router.Handle("/boards/{id:[0-9]+}", handler.Handler(boardHandler.BoardHandler))
	router.Handle("/columns/{id:[0-9]+}", handler.Handler(columnHandler.ColumnHandler))
	router.Handle("/columns", handler.Handler(columnsHandler.ColumnsHandler))
	router.Handle("/cards/positions", handler.Handler(cardPositionsHandler.CardPositionsHandler))
	router.Handle("/cards/{id:[0-9]+}", handler.Handler(cardHandler.CardHandler))
	router.Handle("/cards", handler.Handler(cardsHandler.CardsHandler))
}
