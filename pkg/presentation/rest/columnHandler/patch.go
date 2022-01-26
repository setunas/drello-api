package columnHandler

import (
	"drello-api/pkg/app/usecase/updateColumn"
	"drello-api/pkg/presentation/rest/utils"
	"encoding/json"
	"net/http"
)

func patch(w http.ResponseWriter, r *http.Request, id int) {
	token, err := utils.VerifyIDToken(r.Context(), r)
	if err != nil {
		utils.HandleClientError(w, err, 401, "Invalid token")
		return
	}

	var body struct {
		Title    string
		Position float64
		BoardId  int
	}
	json.NewDecoder(r.Body).Decode(&body)

	column, err := updateColumn.Call(r.Context(), id, body.Title, body.Position, body.BoardId, token.UID)
	if err != nil {
		utils.HandleClientError(w, err, 422, "An error occured during the prosess")
		return
	}

	json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
}
