package rest

import (
	"drello-api/pkg/app/usecase/updateCardPositions"
	"drello-api/pkg/presentation/rest/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func cardPositionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		token, err := utils.VerifyIDToken(r.Context(), r)
		if err != nil {
			utils.HandleClientError(w, err, 401, "Invalid token")
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

		Cards := []updateCardPositions.Card{}
		for _, card := range body.Cards {
			Cards = append(Cards, *updateCardPositions.NewCard(card.ID, card.Position))
		}

		err = updateCardPositions.Call(r.Context(), Cards, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(200)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
