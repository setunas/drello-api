package meHandler

import (
	"drello-api/pkg/presentation/rest/utils"
	"net/http"
)

type meResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		get(w, r)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
