package cardsHandler

import (
	"drello-api/pkg/presentation/rest/util"
	"net/http"
)

type cardResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Position    float64 `json:"position"`
	ColumnId    int     `json:"columnId"`
}

func CardsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		post(w, r)
		return
	}

	util.HandleClientError(w, nil, 404, "Invalid method")
}
