package cardsDS

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (c CardsDS) Update(ctx context.Context, id int, title string, description string, position float64, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE cards SET title = ?, description = ?, position = ?, column_id = ? WHERE id = ?", title, description, position, columnId, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating card: %w", err)
	}

	return domainCard.New(id, title, description, position, columnId), nil
}
