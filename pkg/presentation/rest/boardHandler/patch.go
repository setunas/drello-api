package boardHandler

import (
	"drello-api/pkg/app/usecase/updateBoard"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"encoding/json"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request, id int) error {
	user, err := restutil.AuthenticateUser(r)
	if err != nil {
		return err
	}

	var body struct {
		Title string
	}
	json.NewDecoder(r.Body).Decode(&body)

	ucBoard, err := updateBoard.Call(r.Context(), id, body.Title, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	json.NewEncoder(w).Encode(boardResponse{ID: ucBoard.ID(), Title: ucBoard.Title()})

	return nil
}
