package rest

import (
	"drello-api/pkg/app/usecases/users"
	"encoding/json"
	"net/http"
)

type meResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		token, err := verifyIDToken(r.Context(), r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		output, err := users.GetOne(r.Context(), users.NewGetOneInput(token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(meResponse{
			ID:       output.User.ID(),
			Username: output.User.Username(),
			BoardID:  output.User.BoardID(),
		})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
