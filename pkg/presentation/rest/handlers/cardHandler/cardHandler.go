package cardHandler

import (
	"drello-api/pkg/util/apperr"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type cardResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Position    float64 `json:"position"`
	ColumnId    int     `json:"columnId"`
}

func CardHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return apperr.NewHTTPError(404, "Invalid path ID", nil)
	}

	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodPatch:
		return patch(w, r, id)
	case http.MethodDelete:
		return delete(w, r, id)
	}

	return apperr.NewHTTPError(404, "Invalid method", nil)
}
