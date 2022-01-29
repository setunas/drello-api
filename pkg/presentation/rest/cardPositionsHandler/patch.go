package cardPositionsHandler

import (
	"drello-api/pkg/app/usecase/updateCardPositions"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"encoding/json"
	"fmt"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request) error {
	user, err := util.AuthenticateUser(r)
	if err != nil {
		return err
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

	err = updateCardPositions.Call(r.Context(), Cards, user)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(200)
	return nil
}
