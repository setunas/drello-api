package cardsDS

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (c CardsDS) Create(ctx context.Context, title string, description string, position float64, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO cards (title, description, position, column_id) VALUES (?, ?, ?, ?)", title, description, position, columnId)
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	return domainCard.New(int(lastInsertID), title, description, position, columnId), nil
}
