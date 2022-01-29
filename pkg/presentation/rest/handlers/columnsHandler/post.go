package columnsHandler

import (
	"drello-api/pkg/app/usecase/createColumn"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
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
		Title    string
		Position float64
		BoardId  int
	}
	json.NewDecoder(r.Body).Decode(&body)
	log.Info("Request Body").Add("RequestID", restutil.RetrieveReqID(r.Context())).Add("Body", fmt.Sprintf("%+v", body)).Write()

	column, err := createColumn.Call(r.Context(), body.Title, body.Position, body.BoardId, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
	return nil
}
