package cardsDS

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c CardsDS) Create(ctx context.Context, title string, description string, position float64, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()

	query := "INSERT INTO cards (title, description, position, column_id) VALUES (?, ?, ?, ?)"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("title", title).Add("description", description).
		Add("position", position).Add("columnId", columnId).Write()
	result, err := db.Exec(query, title, description, position, columnId)
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	return domainCard.New(int(lastInsertID), title, description, position, columnId), nil
}
