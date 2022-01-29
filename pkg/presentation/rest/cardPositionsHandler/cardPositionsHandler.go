package cardPositionsHandler

import (
	"drello-api/pkg/util/apperr"
	"net/http"
)

func CardPositionsHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodPatch:
		return patch(w, r)
	}

	return apperr.NewHTTPError(404, "Invalid method", nil)
}
