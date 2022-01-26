package rest

import (
	"drello-api/pkg/app/usecase/deleteColumn"
	"drello-api/pkg/app/usecase/updateColumn"
	"drello-api/pkg/presentation/rest/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type columnResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Position float64 `json:"position"`
	BoardId  int     `json:"boardId"`
}

func columnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		token, err := utils.VerifyIDToken(r.Context(), r)
		if err != nil {
			utils.HandleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Title    string
			Position float64
			BoardId  int
		}
		json.NewDecoder(r.Body).Decode(&body)

		column, err := updateColumn.Call(r.Context(), id, body.Title, body.Position, body.BoardId, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
		return

	case http.MethodDelete:
		token, err := utils.VerifyIDToken(r.Context(), r)
		if err != nil {
			utils.HandleClientError(w, err, 401, "Invalid token")
			return
		}

		err = deleteColumn.Call(r.Context(), id, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
