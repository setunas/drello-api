package cardHandler

import (
	"drello-api/pkg/presentation/rest/util"
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

func CardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.HandleClientError(w, err, 400, "Invalid ID.")
		return
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

	util.HandleClientError(w, nil, 404, "Invalid method")
}
