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
	switch r.Method {
	case http.MethodPost:
		boardId, err := strconv.Atoi(r.FormValue("boardId"))
		if err != nil {
			handleClientError(w, err, 400, "Invalid boardId.")
			return
		}

		output, err := columns.Create(r.Context(), datasource.Column{}, columns.NewCreateInput(r.FormValue("title"), boardId))
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodPatch:
		boardId, err := strconv.Atoi(r.FormValue("boardId"))
		if err != nil {
			handleClientError(w, err, 400, "Invalid boardId.")
			return
		}

		output, err := columns.Update(r.Context(), datasource.Column{}, columns.NewUpdateInput(id, r.FormValue("title"), boardId))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(columnResponse{ID: output.Column.ID(), Title: output.Column.Title(), BoardId: output.Column.BoardId()})
		return

	case http.MethodDelete:
		err = columns.Delete(r.Context(), datasource.Column{}, columns.NewDeleteInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
