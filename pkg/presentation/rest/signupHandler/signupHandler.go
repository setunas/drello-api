package signupHandler

import (
	"drello-api/pkg/presentation/rest/utils"
	"net/http"
)

type signupResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		post(w, r)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
