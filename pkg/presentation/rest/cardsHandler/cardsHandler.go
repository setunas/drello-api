package cardsHandler

import (
	"drello-api/pkg/util/apperr"
	"net/http"
)

type cardResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Position    float64 `json:"position"`
	ColumnId    int     `json:"columnId"`
}

func CardsHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodOptions:
		return nil
	case http.MethodPost:
		return post(w, r)
	}

	return apperr.NewHTTPError(404, "Invalid method", nil)
}
