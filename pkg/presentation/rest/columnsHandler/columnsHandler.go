package columnsHandler

import (
	"drello-api/pkg/presentation/rest/utils"
	"net/http"
)

type columnResponse struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Position float64 `json:"position"`
	BoardId  int     `json:"boardId"`
}

func ColumnsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		post(w, r)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
