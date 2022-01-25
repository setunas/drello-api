package rest

import (
	"drello-api/pkg/app/usecase/updateCardPositions"
	"encoding/json"
	"fmt"
	"net/http"
)

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

		inputCards := []updateCardPositions.InputCard{}
		for _, card := range body.Cards {
			inputCards = append(inputCards, *updateCardPositions.NewInputCard(card.ID, card.Position))
		}

		err = updateCardPositions.UpdateCardPositions(r.Context(), inputCards, token.UID)
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(200)
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
