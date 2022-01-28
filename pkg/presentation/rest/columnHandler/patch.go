package columnHandler

import (
	"drello-api/pkg/app/usecase/updateColumn"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"encoding/json"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request, id int) error {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		return myerr.NewHTTPError(401, "Invalid token", err)
	}

	var body struct {
		Title    string
		Position float64
		BoardId  int
	}
	json.NewDecoder(r.Body).Decode(&body)

	column, err := updateColumn.Call(r.Context(), id, body.Title, body.Position, body.BoardId, token.UID)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
	return nil
}
