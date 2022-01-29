package cardsDS

import (
	"context"
	"database/sql"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c CardsDS) GetOneByID(ctx context.Context, id int) (*domainCard.Card, error) {
	var title string
	var description string
	var position float64
	var columnID int

	db := mysql.DBPool()
	query := "SELECT title, description, position, column_id FROM cards WHERE id = ?"
	log.Info("SQL").Add("SQL", query).Add("id", id).Write()
	row := db.QueryRow(query, id)

	switch err := row.Scan(&title, &description, &position, &columnID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainCard.New(id, title, description, position, columnID), nil
	default:
		return nil, err
	}
}
