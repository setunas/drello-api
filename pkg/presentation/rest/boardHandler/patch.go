package boardHandler

import (
	"drello-api/pkg/app/usecase/updateBoard"
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
		Title string
	}
	json.NewDecoder(r.Body).Decode(&body)

	ucBoard, err := updateBoard.Call(r.Context(), id, body.Title, token.UID)
	if err != nil {
		utils.HandleClientError(w, err, 422, "An error occured during the prosess")
		return
	}

	json.NewEncoder(w).Encode(boardResponse{ID: ucBoard.ID(), Title: ucBoard.Title()})
}
