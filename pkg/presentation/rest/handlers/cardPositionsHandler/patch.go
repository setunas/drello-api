package cardPositionsHandler

import (
	"drello-api/pkg/app/usecase/updateCardPositions"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"encoding/json"
	"fmt"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request) error {
	user, err := restutil.AuthenticateUser(r)
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
	log.Info("Request Body").Add("RequestID", restutil.RetrieveReqID(r.Context())).Add("Body", fmt.Sprintf("%+v", body)).Write()

	Cards := []updateCardPositions.Card{}
	for _, card := range body.Cards {
		Cards = append(Cards, *updateCardPositions.NewCard(card.ID, card.Position))
	}

	err = updateCardPositions.Call(r.Context(), Cards, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(200)
	return nil
}
