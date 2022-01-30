package handler

import (
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"encoding/json"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
