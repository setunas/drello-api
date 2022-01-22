package rest

import (
	"drello-api/pkg/app/cards"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"fmt"
	"net/http"
)

type positionResponse struct {
	Cards []cardResponse `json:"cards"`
}

func positionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		token, err := verifyIDToken(r.Context(), r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		type card struct {
			ID       int
			Position float64
		}

		var body struct {
			Cards []card
		}
		json.NewDecoder(r.Body).Decode(&body)
		fmt.Println("body", body)

		output, err := cards.Update(r.Context(), datasource.Column{}, datasource.Card{}, datasource.User{}, cards.NewUpdateInput(id, body.Title, body.Description, body.Position, body.ColumnID, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(cardResponse{ID: output.Card.ID(), Title: output.Card.Title(), Description: output.Card.Description(), Position: output.Card.Position(), ColumnId: output.Card.ColumnId()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
