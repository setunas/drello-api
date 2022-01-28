package boardHandler

import (
	"drello-api/pkg/util/myerr"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BoardHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return myerr.NewHTTPError(404, "Invalid path ID", nil)
	}

	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodGet:
		get(w, r, id)
		return nil
	case http.MethodPatch:
		patch(w, r, id)
		return nil
	}

	return myerr.NewHTTPError(404, "Invalid method", nil)
}

type boardResponse struct {
	ID      int              `json:"id"`
	Title   string           `json:"title"`
	Columns []columnResponse `json:"columns"`
	Cards   []cardResponse   `json:"cards"`
}

type columnResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Position float64 `json:"position"`
	BoardId  int     `json:"boardId"`
}

type cardResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Position    float64 `json:"position"`
	ColumnId    int     `json:"columnId"`
}
