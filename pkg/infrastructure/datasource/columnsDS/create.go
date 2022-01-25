package columnsDS

import (
	"context"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (c ColumnsDS) Create(ctx context.Context, title string, position float64, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO columns (title, position, board_id) VALUES (?, ?, ?)", title, position, boardId)
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	return domainColumn.New(int(lastInsertID), title, position, boardId), nil
}