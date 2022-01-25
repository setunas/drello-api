package boardsDS

import (
	"context"
	"database/sql"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (b BoardsDS) GetOne(ctx context.Context, id int) (*domainBoard.Board, error) {
	var title string

	db := mysql.DBPool()
	row := db.QueryRow("SELECT title FROM boards WHERE id = ?", id)

	switch err := row.Scan(&title); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainBoard.New(id, title), nil
	default:
		return nil, err
	}
}
