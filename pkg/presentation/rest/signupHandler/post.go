package signupHandler

import (
	"drello-api/pkg/app/usecase/signup"
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
		Username string
		Title    string
	}
	json.NewDecoder(r.Body).Decode(&body)

	user, err := signup.Call(r.Context(), body.Username, token.UID, body.Title)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signupResponse{ID: user.ID(), Username: user.Username(), BoardID: user.BoardID()})
	return nil
}
