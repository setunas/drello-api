package rest

import (
	"drello-api/pkg/app/usecase/deleteCard"
	"drello-api/pkg/app/usecase/updateCard"
	"drello-api/pkg/presentation/rest/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type cardResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Position    float64 `json:"position"`
	ColumnId    int     `json:"columnId"`
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleClientError(w, err, 400, "Invalid ID.")
		return
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
			Title       string
			Description string
			Position    float64
			ColumnID    int
		}
		json.NewDecoder(r.Body).Decode(&body)
		fmt.Println("body", body)

		ucCard, err := updateCard.Call(r.Context(), id, body.Title, body.Description, body.Position, body.ColumnID, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(cardResponse{ID: ucCard.ID(), Title: ucCard.Title(), Description: ucCard.Description(), Position: ucCard.Position(), ColumnId: ucCard.ColumnId()})
		return

	case http.MethodDelete:
		token, err := utils.VerifyIDToken(r.Context(), r)
		if err != nil {
			utils.HandleClientError(w, err, 401, "Invalid token")
			return
		}

		err = deleteCard.Call(r.Context(), id, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
