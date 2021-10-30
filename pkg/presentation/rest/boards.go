package rest

import (
	"drello-api/pkg/app/boards"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type boardResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodGet:
		output, err := boards.GetOne(r.Context(), datasource.Board{}, boards.NewGetOneInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(&boardResponse{
			ID:    output.Board.ID(),
			Title: output.Board.Title(),
		})
		return

	case http.MethodPatch:
		output, err := boards.Update(r.Context(), datasource.Board{}, boards.NewUpdateInput(id, r.FormValue("title")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(boardResponse{ID: output.Board.ID(), Title: output.Board.Title()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
