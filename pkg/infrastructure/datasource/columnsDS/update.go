package columnsDS

import (
	"context"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (c ColumnsDS) Update(ctx context.Context, id int, title string, position float64, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE columns SET title = ?, position = ? WHERE id = ?", title, position, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating column: %w", err)
	}

	return domainColumn.New(id, title, position, boardId), nil
}