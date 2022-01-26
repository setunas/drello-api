package signupHandler

import (
	"drello-api/pkg/app/usecase/signup"
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
		Username string
		Title    string
	}
	json.NewDecoder(r.Body).Decode(&body)

	user, err := signup.Call(r.Context(), body.Username, token.UID, body.Title)
	if err != nil {
		utils.HandleClientError(w, err, 422, "An error occured while creating a user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signupResponse{ID: user.ID(), Username: user.Username(), BoardID: user.BoardID()})
}
