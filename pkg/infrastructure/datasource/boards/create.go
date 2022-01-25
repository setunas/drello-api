package boards

import (
	"context"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (w Board) Create(ctx context.Context, title string) (*domainBoard.Board, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO boards (title) VALUES (?)", title)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}

	return domainBoard.New(int(lastInsertID), title), nil
}
