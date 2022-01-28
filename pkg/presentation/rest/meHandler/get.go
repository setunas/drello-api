package meHandler

import (
	"drello-api/pkg/app/usecase/getMe"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) error {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		return myerr.NewHTTPError(401, "Invalid token", err)
	}

	user, err := getMe.Call(r.Context(), token.UID)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	json.NewEncoder(w).Encode(meResponse{
		ID:       user.ID(),
		Username: user.Username(),
		BoardID:  user.BoardID(),
	})
	return nil
}
