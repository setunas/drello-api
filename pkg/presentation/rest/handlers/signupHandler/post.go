package signupHandler

import (
	"drello-api/pkg/app/usecase/signup"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"encoding/json"
	"fmt"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) error {
	token, err := restutil.VerifyIDToken(r.Context(), r)
	if err != nil {
		return apperr.NewHTTPError(401, "Invalid token", err)
	}

	var body struct {
		Username string
		Title    string
	}
	json.NewDecoder(r.Body).Decode(&body)
	log.Info("Request Body").Add("RequestID", restutil.RetrieveReqID(r.Context())).Add("Body", fmt.Sprintf("%+v", body)).Write()

	user, err := signup.Call(r.Context(), body.Username, token.UID, body.Title)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signupResponse{ID: user.ID(), Username: user.Username(), BoardID: user.BoardID()})
	return nil
}
