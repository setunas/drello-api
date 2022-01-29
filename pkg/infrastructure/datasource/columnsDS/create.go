package columnsDS

import (
	"context"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c ColumnsDS) Create(ctx context.Context, title string, position float64, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()

	query := "INSERT INTO columns (title, position, board_id) VALUES (?, ?, ?)"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("title", title).
		Add("position", position).Add("boardId", boardId).Write()
	result, err := db.Exec(query, title, position, boardId)
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	return domainColumn.New(int(lastInsertID), title, position, boardId), nil
}
