package columnsHandler

import (
	"drello-api/pkg/app/usecase/createColumn"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"encoding/json"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) error {
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

	column, err := createColumn.Call(r.Context(), body.Title, body.Position, body.BoardId, token.UID)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
	return nil
}
