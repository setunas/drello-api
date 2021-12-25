package rest

import (
	"drello-api/pkg/app/cards"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type cardResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ColumnId    int    `json:"columnId"`
}

func cardsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		token, err := verifyIDToken(r.Context(), r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Title       string
			Description string
			ColumnID    int
		}
		json.NewDecoder(r.Body).Decode(&body)

		output, err := cards.Create(r.Context(), datasource.Column{}, datasource.Card{}, datasource.User{}, cards.NewCreateInput(body.Title, body.Description, body.ColumnID, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cardResponse{ID: output.Card.ID(), Title: output.Card.Title(), Description: output.Card.Description(), ColumnId: output.Card.ColumnId()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
		return
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		var body struct {
			Title       string
			Description string
			ColumnID    int
		}
		json.NewDecoder(r.Body).Decode(&body)

		output, err := cards.Update(r.Context(), datasource.Card{}, cards.NewUpdateInput(id, body.Title, body.Description, body.ColumnID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(cardResponse{ID: output.Card.ID(), Title: output.Card.Title(), Description: output.Card.Description(), ColumnId: output.Card.ColumnId()})
		return

	case http.MethodDelete:
		err = cards.Delete(r.Context(), datasource.Card{}, cards.NewDeleteInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
