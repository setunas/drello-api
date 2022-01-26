package cardHandler

import (
	"drello-api/pkg/app/usecase/deleteCard"
	"drello-api/pkg/presentation/rest/utils"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request, id int) {
	token, err := utils.VerifyIDToken(r.Context(), r)
	if err != nil {
		utils.HandleClientError(w, err, 401, "Invalid token")
		return
	}

	err = deleteCard.Call(r.Context(), id, token.UID)
	if err != nil {
		utils.HandleClientError(w, err, 422, "An error occured during the prosess")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
