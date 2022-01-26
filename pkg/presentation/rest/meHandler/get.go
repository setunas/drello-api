package meHandler

import (
	"drello-api/pkg/app/usecase/getMe"
	"drello-api/pkg/presentation/rest/util"
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		util.HandleClientError(w, err, 401, "Invalid token")
		return
	}

	user, err := getMe.Call(r.Context(), token.UID)
	if err != nil {
		util.HandleClientError(w, err, 422, "An error occured during the prosess")
		return
	}

	json.NewEncoder(w).Encode(meResponse{
		ID:       user.ID(),
		Username: user.Username(),
		BoardID:  user.BoardID(),
	})
}
