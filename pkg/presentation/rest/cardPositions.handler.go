package rest

import (
	"drello-api/pkg/app/usecases/cards"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"fmt"
	"net/http"
)

type positionResponse struct {
	Cards []cardResponse `json:"cards"`
}

func cardPositionsHandler(w http.ResponseWriter, r *http.Request) {
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

		cds := []cards.CardInput{}
		for _, card := range body.Cards {
			cds = append(cds, *cards.NewCardInput(card.ID, card.Position))
		}

		err = cards.UpdatePositions(r.Context(), datasource.Column{}, datasource.Card{}, datasource.User{}, cards.NewUpdatePositionsInput(cds, token.UID))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(200)
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
