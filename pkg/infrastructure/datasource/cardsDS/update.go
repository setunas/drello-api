package cardsDS

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c CardsDS) Update(ctx context.Context, id int, title string, description string, position float64, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()
	query := "UPDATE cards SET title = ?, description = ?, position = ?, column_id = ? WHERE id = ?"
	log.Info("SQL").Add("SQL", query).Add("title", title).Add("description", description).
		Add("position", position).Add("columnId", columnId).Write()
	_, err := db.Exec(query, title, description, position, columnId, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating card: %w", err)
	}

	return domainCard.New(id, title, description, position, columnId), nil
}
