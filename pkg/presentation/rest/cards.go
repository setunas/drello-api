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
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func cardsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		output, err := cards.Create(r.Context(), datasource.Card{}, cards.NewCreateInput(r.FormValue("title"), r.FormValue("description")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cardResponse{ID: output.Card.ID(), Title: output.Card.Title()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodPatch:
		output, err := cards.Update(r.Context(), datasource.Card{}, cards.NewUpdateInput(id, r.FormValue("title"), r.FormValue("description")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(cardResponse{ID: output.Card.ID(), Title: output.Card.Title()})
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
