package rest

import (
	"drello-api/pkg/app/boards"
	"drello-api/pkg/app/users"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
)

type usreResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

		boardOutput, err := boards.Create(r.Context(), datasource.Board{}, boards.NewCreateInput(r.FormValue("username")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured while creating a board for the user")
			return
		}

		userOutput, err := users.Create(r.Context(), datasource.User{}, users.NewCreateInput(r.FormValue("username"), boardOutput.Board.ID(), token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured while creating a user")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(usreResponse{ID: userOutput.User.ID(), Username: userOutput.User.Username(), BoardID: userOutput.User.BoardID()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
