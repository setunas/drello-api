package cardsDS

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/log"
	"fmt"
	"strings"
)

func (c CardsDS) GetListByIDs(ctx context.Context, ids []int) ([]*domainCard.Card, error) {
	db := mysql.DBPool()

	cards := []*domainCard.Card{}

	if len(ids) == 0 {
		return cards, nil
	}

	args := make([]interface{}, len(ids))
	for i := range ids {
		args[i] = ids[i]
	}

	query := "SELECT id, title, description, position, column_id FROM cards WHERE id IN (?" + strings.Repeat(",?", len(ids)-1) + ")"
	log.Info("SQL").Add("SQL", query).Add("args...", args).Write()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed querying cards: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var description string
		var position float64
		var columnId int

		err := rows.Scan(&id, &title, &description, &position, &columnId)
		if err != nil {
			return nil, fmt.Errorf("failed scanning card rows: %w", err)
		}

		cards = append(cards, domainCard.New(id, title, description, position, columnId))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying cards: %w", err)
	}

	return cards, nil
}
