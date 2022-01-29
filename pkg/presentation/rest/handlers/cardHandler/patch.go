package cardHandler

import (
	"drello-api/pkg/app/usecase/updateCard"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"encoding/json"
	"fmt"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request, id int) error {
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
	log.Info("Request Body").Add("RequestID", restutil.RetrieveReqID(r.Context())).Add("Body", fmt.Sprintf("%+v", body)).Write()

	ucCard, err := updateCard.Call(r.Context(), id, body.Title, body.Description, body.Position, body.ColumnID, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	json.NewEncoder(w).Encode(cardResponse{ID: ucCard.ID(), Title: ucCard.Title(), Description: ucCard.Description(), Position: ucCard.Position(), ColumnId: ucCard.ColumnId()})
	return nil
}
