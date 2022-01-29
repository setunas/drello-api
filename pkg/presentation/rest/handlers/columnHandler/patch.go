package columnHandler

import (
	"drello-api/pkg/app/usecase/updateColumn"
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
		Title    string
		Position float64
		BoardId  int
	}
	json.NewDecoder(r.Body).Decode(&body)
	log.Info("Request Body").Add("RequestID", restutil.RetrieveReqID(r.Context())).Add("Body", fmt.Sprintf("%+v", body)).Write()

	column, err := updateColumn.Call(r.Context(), id, body.Title, body.Position, body.BoardId, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
	return nil
}
