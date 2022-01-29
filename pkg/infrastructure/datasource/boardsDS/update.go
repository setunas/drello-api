package boardsDS

import (
	"context"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/log"
	"fmt"
)

func (b BoardsDS) Update(ctx context.Context, id int, title string) (*domainBoard.Board, error) {
	db := mysql.DBPool()

	query := "UPDATE boards SET title = ? WHERE id = ?"
	log.Info("SQL").Add("SQL", query).Add("title", title).Add("id", id).Write()

	_, err := db.Exec(query, title, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating board: %w", err)
	}

	return domainBoard.New(id, title), nil
}
