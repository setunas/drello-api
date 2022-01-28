package cardPositionsHandler

import (
	"drello-api/pkg/util/myerr"
	"net/http"
)

func CardPositionsHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodOptions:
		return nil

	case http.MethodPatch:
		patch(w, r)
		return nil
	}

	return myerr.NewHTTPError(404, "Invalid method", nil)
}
