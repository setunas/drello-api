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
		columnId, err := strconv.Atoi(r.FormValue("columnId"))
		if err != nil {
			handleClientError(w, err, 400, "Invalid columnId.")
			return
		}

		output, err := cards.Create(r.Context(), datasource.Card{}, cards.NewCreateInput(r.FormValue("title"), r.FormValue("description"), columnId))
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
		columnId, err := strconv.Atoi(r.FormValue("columnId"))
		if err != nil {
			handleClientError(w, err, 400, "Invalid columnId.")
		}

		output, err := cards.Update(r.Context(), datasource.Card{}, cards.NewUpdateInput(id, r.FormValue("title"), r.FormValue("description"), columnId))
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
