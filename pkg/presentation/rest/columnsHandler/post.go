package columnsHandler

import (
	"drello-api/pkg/app/usecase/createColumn"
	"drello-api/pkg/presentation/rest/utils"
	"encoding/json"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {
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

	column, err := createColumn.Call(r.Context(), body.Title, body.Position, body.BoardId, token.UID)
	if err != nil {
		utils.HandleClientError(w, err, 422, "An error occured during the prosess")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(columnResponse{ID: column.ID(), Title: column.Title(), Position: column.Positon(), BoardId: column.BoardId()})
}
