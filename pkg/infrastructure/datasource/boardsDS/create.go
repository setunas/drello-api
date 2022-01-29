package boardsDS

import (
	"context"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (b BoardsDS) Create(ctx context.Context, title string) (*domainBoard.Board, error) {
	db := mysql.DBPool()

	query := "INSERT INTO boards (title) VALUES (?)"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("title", title).Write()

	result, err := db.Exec(query, title)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}

	return domainBoard.New(int(lastInsertID), title), nil
}
