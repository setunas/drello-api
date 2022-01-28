package cardsHandler

import (
	"drello-api/pkg/app/usecase/createCard"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"encoding/json"
	"fmt"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) error {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		return myerr.NewHTTPError(401, "Invalid token", err)
	}

	var body struct {
		Title       string
		Description string
		Position    float64
		ColumnID    int
	}
	json.NewDecoder(r.Body).Decode(&body)
	fmt.Println("body", body)

	ucCard, err := createCard.Call(r.Context(), body.Title, body.Description, body.Position, body.ColumnID, token.UID)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cardResponse{ID: ucCard.ID(), Title: ucCard.Title(), Description: ucCard.Description(), Position: ucCard.Position(), ColumnId: ucCard.ColumnId()})
	return nil
}
