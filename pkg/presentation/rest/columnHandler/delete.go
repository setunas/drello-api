package columnHandler

import (
	"drello-api/pkg/app/usecase/deleteColumn"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request, id int) error {
	user, err := restutil.AuthenticateUser(r)
	if err != nil {
		return err
	}

	err = deleteColumn.Call(r.Context(), id, user)
	if err != nil {
		return apperr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
