package columnHandler

import (
	"drello-api/pkg/app/usecase/deleteColumn"
	"drello-api/pkg/presentation/rest/util"
	"drello-api/pkg/util/myerr"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request, id int) error {
	user, err := util.AuthenticateUser(r)
	if err != nil {
		return err
	}

	err = deleteColumn.Call(r.Context(), id, user)
	if err != nil {
		return myerr.NewHTTPError(500, "An error occured during the prosess", err)
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
