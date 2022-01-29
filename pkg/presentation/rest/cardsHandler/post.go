package cardsHandler

import (
	"drello-api/pkg/app/usecase/createCard"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"encoding/json"
	"fmt"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) error {
	user, err := restutil.AuthenticateUser(r)
	if err != nil {
		return err
	}

	var body struct {
		Title       string
		Description string
		Position    float64
		ColumnID    int
	}
	json.NewDecoder(r.Body).Decode(&body)
	fmt.Println("body", body)

	ucCard, err := createCard.Call(r.Context(), body.Title, body.Description, body.Position, body.ColumnID, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cardResponse{ID: ucCard.ID(), Title: ucCard.Title(), Description: ucCard.Description(), Position: ucCard.Position(), ColumnId: ucCard.ColumnId()})
	return nil
}
