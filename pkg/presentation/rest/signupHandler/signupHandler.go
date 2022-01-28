package signupHandler

import (
	"drello-api/pkg/util/myerr"
	"net/http"
)

type signupResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodOptions:
		return nil

	case http.MethodPost:
		post(w, r)
		return nil
	}

	return myerr.NewHTTPError(404, "Invalid method", nil)
}
