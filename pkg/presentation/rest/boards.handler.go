package rest

import (
	"drello-api/pkg/app/usecase/getBoardWithColumnsAndCards"
	"drello-api/pkg/app/usecase/updateBoard"
	"drello-api/pkg/presentation/rest/utils"
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.HandleClientError(w, err, 400, "Invalid ID.")
		return
	}

	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodGet:
		token, err := utils.VerifyIDToken(r.Context(), r)
		if err != nil {
			utils.HandleClientError(w, err, 401, "Invalid token")
			return
		}
		ucBoard, ucColumns, ucCards, err := getBoardWithColumnsAndCards.Call(r.Context(), id, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		columns := []columnResponse{}
		cards := []cardResponse{}

		for _, column := range ucColumns {
			columns = append(columns, columnResponse{
				ID:       column.ID(),
				Title:    column.Title(),
				Position: column.Positon(),
				BoardId:  column.BoardId(),
			})
		}
		for _, card := range ucCards {
			cards = append(cards, cardResponse{
				ID:          card.ID(),
				Title:       card.Title(),
				Description: card.Description(),
				Position:    card.Position(),
				ColumnId:    card.ColumnId(),
			})
		}

		json.NewEncoder(w).Encode(boardResponse{
			ID:      ucBoard.ID(),
			Title:   ucBoard.Title(),
			Columns: columns,
			Cards:   cards,
		})
		return

	case http.MethodPatch:
		token, err := utils.VerifyIDToken(r.Context(), r)
		if err != nil {
			utils.HandleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Title string
		}
		json.NewDecoder(r.Body).Decode(&body)

		ucBoard, err := updateBoard.Call(r.Context(), id, body.Title, token.UID)
		if err != nil {
			utils.HandleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(boardResponse{ID: ucBoard.ID(), Title: ucBoard.Title()})
		return
	}

	utils.HandleClientError(w, nil, 404, "Invalid method")
}
