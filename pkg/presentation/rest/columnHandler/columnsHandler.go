package columnHandler

import (
	"drello-api/pkg/util/myerr"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type columnResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Position float64 `json:"position"`
	BoardId  int     `json:"boardId"`
}

func ColumnHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return myerr.NewHTTPError(404, "Invalid path ID", nil)
	}

	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodPatch:
		return patch(w, r, id)
	case http.MethodDelete:
		return delete(w, r, id)
	}

	return myerr.NewHTTPError(404, "Invalid method", nil)
}
