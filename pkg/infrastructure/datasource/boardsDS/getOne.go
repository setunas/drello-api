package boardsDS

import (
	"context"
	"database/sql"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"fmt"
)

func (b BoardsDS) GetOne(ctx context.Context, id int) (*domainBoard.Board, error) {
	var title string

	db := mysql.DBPool()
	query := "SELECT title FROM boards WHERE id = ?"
	log.Info("SQL").Add("SQL", query).Add("id", id).Write()
	row := db.QueryRow(query, id)

	switch err := row.Scan(&title); err {
	case sql.ErrNoRows:
		return nil,
			apperr.NewAppError(
				[]apperr.Tag{apperr.RecordNotFound},
				fmt.Sprintf("board record not found with id %d", id),
				nil,
			)
	case nil:
		return domainBoard.New(id, title), nil
	default:
		return nil, err
	}
}
