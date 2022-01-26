package boardHandler

import (
	"drello-api/pkg/app/usecase/getBoardWithColumnsAndCards"
	"drello-api/pkg/presentation/rest/util"
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request, id int) {
	token, err := util.VerifyIDToken(r.Context(), r)
	if err != nil {
		util.HandleClientError(w, err, 401, "Invalid token")
		return
	}
	ucBoard, ucColumns, ucCards, err := getBoardWithColumnsAndCards.Call(r.Context(), id, token.UID)
	if err != nil {
		util.HandleClientError(w, err, 422, "An error occured during the prosess")
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
}
