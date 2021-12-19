package rest

import (
	"drello-api/pkg/app/boards"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type boardResponse struct {
	ID      int              `json:"id"`
	Title   string           `json:"title"`
	Columns []columnResponse `json:"columns"`
	Cards   []cardResponse   `json:"cards"`
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
		return
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		// verifyIDToken(ctx, r)
		output, err := boards.GetOne(ctx, datasource.Board{}, datasource.Column{}, datasource.Card{}, boards.NewGetOneInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		columns := []columnResponse{}
		cards := []cardResponse{}

		for _, column := range output.Columns {
			columns = append(columns, columnResponse{
				ID:      column.ID(),
				Title:   column.Title(),
				BoardId: column.BoardId(),
			})
		}
		for _, card := range output.Cards {
			cards = append(cards, cardResponse{
				ID:          card.ID(),
				Title:       card.Title(),
				Description: card.Description(),
				ColumnId:    card.ColumnId(),
			})
		}

		json.NewEncoder(w).Encode(boardResponse{
			ID:      output.Board.ID(),
			Title:   output.Board.Title(),
			Columns: columns,
			Cards:   cards,
		})
		return

	case http.MethodPatch:
		output, err := boards.Update(ctx, datasource.Board{}, boards.NewUpdateInput(id, r.FormValue("title")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(boardResponse{ID: output.Board.ID(), Title: output.Board.Title()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}