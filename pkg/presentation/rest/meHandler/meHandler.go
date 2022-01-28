package meHandler

import (
	"drello-api/pkg/util/myerr"
	"net/http"
)

type meResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func MeHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodGet:
		return get(w, r)
	}

	return myerr.NewHTTPError(404, "Invalid method", nil)
}
