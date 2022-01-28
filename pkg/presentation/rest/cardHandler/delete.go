package cardHandler

import (
	"drello-api/pkg/app/usecase/deleteCard"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request, id int) error {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		return myerr.NewHTTPError(401, "Invalid token", err)
	}

	err = deleteCard.Call(r.Context(), id, token.UID)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
