package rest

import (
	"drello-api/pkg/app/columns"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type columnResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	BoardId int    `json:"boardId"`
}

func columnsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		token, err := verifyIDToken(ctx, r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Title    string
			Position float64
			BoardId  int
		}
		json.NewDecoder(r.Body).Decode(&body)

		output, err := columns.Create(ctx, datasource.Column{}, datasource.User{}, columns.NewCreateInput(body.Title, body.Position, body.BoardId, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(columnResponse{ID: output.Column.ID(), Title: output.Column.Title(), BoardId: output.Column.BoardId()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}

func columnHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		token, err := verifyIDToken(ctx, r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Title    string
			Position float64
			BoardId  int
		}
		json.NewDecoder(r.Body).Decode(&body)

		output, err := columns.Update(ctx, datasource.Column{}, datasource.User{}, columns.NewUpdateInput(id, body.Title, body.Position, body.BoardId, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(columnResponse{ID: output.Column.ID(), Title: output.Column.Title(), BoardId: output.Column.BoardId()})
		return

	case http.MethodDelete:
		token, err := verifyIDToken(ctx, r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		err = columns.Delete(ctx, datasource.Board{}, datasource.Column{}, datasource.User{}, columns.NewDeleteInput(id, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
