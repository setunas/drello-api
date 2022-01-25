package boards

import (
	"context"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (w Board) Update(ctx context.Context, id int, title string) (*domainBoard.Board, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE boards SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating board: %w", err)
	}

	return domainBoard.New(id, title), nil
}
