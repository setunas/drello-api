package columnHandler

import (
	"drello-api/pkg/presentation/rest/utils"
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

func ColumnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPatch:
		patch(w, r, id)
		return

	case http.MethodDelete:
		delete(w, r, id)
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
