package rest

import (
	"drello-api/pkg/app/users"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usreResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
		return
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		token, err := verifyIDToken(ctx, r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		output, err := users.GetOne(ctx, datasource.User{}, users.NewGetOneInput(token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(usreResponse{
			ID:       output.User.ID(),
			Username: output.User.Username(),
			BoardID:  output.User.BoardID(),
		})
		return

	case http.MethodPost:
		token, err := verifyIDToken(ctx, r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		boardID, err := strconv.Atoi(r.FormValue("boardId"))
		if err != nil {
			handleClientError(w, err, 400, "Invalid boardId.")
			return
		}

		output, err := users.Create(r.Context(), datasource.User{}, users.NewCreateInput(r.FormValue("username"), boardID, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(usreResponse{ID: output.User.ID(), Username: output.User.Username(), BoardID: output.User.BoardID()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
