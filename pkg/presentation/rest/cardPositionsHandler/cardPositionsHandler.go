package cardPositionsHandler

import (
	"drello-api/pkg/presentation/rest/util"
	"net/http"
)

func CardPositionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		patch(w, r)
		return
	}

	util.HandleClientError(w, nil, 404, "Invalid method")
}
