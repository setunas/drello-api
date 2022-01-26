package cardHandler

import (
	"drello-api/pkg/app/usecase/deleteCard"
	"drello-api/pkg/presentation/rest/util"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request, id int) {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		util.HandleClientError(w, err, 401, "Invalid token")
		return
	}

	err = deleteCard.Call(r.Context(), id, token.UID)
	if err != nil {
		util.HandleClientError(w, err, 422, "An error occured during the prosess")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
