package rest

import (
	"drello-api/pkg/presentation/rest/handlers/boardHandler"
	"drello-api/pkg/presentation/rest/handlers/cardHandler"
	"drello-api/pkg/presentation/rest/handlers/cardPositionsHandler"
	"drello-api/pkg/presentation/rest/handlers/cardsHandler"
	"drello-api/pkg/presentation/rest/handlers/columnHandler"
	"drello-api/pkg/presentation/rest/handlers/columnsHandler"
	"drello-api/pkg/presentation/rest/handlers/meHandler"
	"drello-api/pkg/presentation/rest/handlers/signupHandler"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"encoding/json"
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
	r = restutil.AppendReqIDToCtx(r)
	logHTTPRequest(r)
	setHeaders(w)

	err := fn(w, r)
	if err == nil {
		return
	}

	handleError(w, r, err)
}

func logHTTPRequest(r *http.Request) {
	log.Info("Got a HTTP Request").
		Add("RequestID", restutil.RetrieveReqID(r.Context())).
		Add("Method", r.Method).
		Add("URI", r.URL.String()).
		Add("Referer", r.Header.Get("Referer")).
		Add("UserAgent", r.Header.Get("User-Agent")).
		Write()
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	reqID := restutil.RetrieveReqID(r.Context())

	type errRes struct {
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")

	httpError, ok := err.(*apperr.HTTPError)
	if !ok {
		log.Errorf("Unknown error type: %v", err).Add("RequestID", reqID).Write()
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(errRes{Message: "Unknown error occured"})
		return
	}

	log.Error("HTTP Error: ", err).Add("RequestID", reqID).Add("OccurredAt", httpError.OccurredAt()).Write()
	w.WriteHeader(httpError.Status())
	if httpError.IsClientError() {
		json.NewEncoder(w).Encode(errRes{Message: httpError.Error()})
	} else {
		json.NewEncoder(w).Encode(errRes{Message: "Internal server error"})
	}
}
