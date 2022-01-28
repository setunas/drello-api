package boardsDS

import (
	"context"
	"database/sql"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/myerr"
	"fmt"
)

func (b BoardsDS) GetOne(ctx context.Context, id int) (*domainBoard.Board, error) {
	var title string

	db := mysql.DBPool()
	row := db.QueryRow("SELECT title FROM boards WHERE id = ?", id)

	switch err := row.Scan(&title); err {
	case sql.ErrNoRows:
		return nil, myerr.NewHTTPError(404, fmt.Sprintf("board record not found with id %d", id), nil)
	case nil:
		return domainBoard.New(id, title), nil
	default:
		return nil, err
	}
}
