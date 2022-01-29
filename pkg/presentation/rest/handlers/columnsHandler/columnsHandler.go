package columnsHandler

import (
	"drello-api/pkg/util/apperr"
	"net/http"
)

type columnResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Position float64 `json:"position"`
	BoardId  int     `json:"boardId"`
}

func ColumnsHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodPost:
		return post(w, r)
	}

	return apperr.NewHTTPError(404, "Invalid method", nil)
}
