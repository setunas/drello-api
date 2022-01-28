package cardHandler

import (
	"drello-api/pkg/app/usecase/updateCard"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"encoding/json"
	"fmt"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request, id int) error {
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

	ucCard, err := updateCard.Call(r.Context(), id, body.Title, body.Description, body.Position, body.ColumnID, token.UID)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	json.NewEncoder(w).Encode(cardResponse{ID: ucCard.ID(), Title: ucCard.Title(), Description: ucCard.Description(), Position: ucCard.Position(), ColumnId: ucCard.ColumnId()})
	return nil
}
