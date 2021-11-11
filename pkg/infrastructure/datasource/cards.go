package datasource

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
	"strings"
)

type Card struct{}

func (c Card) GetListByColumnIds(ctx context.Context, columnIds []int) (*[]*domainCard.Card, error) {
	db := mysql.DBPool()

	var args []interface{}
	for _, v := range columnIds {
		args = append(args, v)
	}

	sql := "SELECT id, title, description, column_id FROM cards WHERE column_id IN (?" + strings.Repeat(",?", len(columnIds)-1) + ")"

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed querying cards: %w", err)
	}
	defer rows.Close()

	cards := []*domainCard.Card{}
	for rows.Next() {
		var id int
		var title string
		var description string
		var columnId int

		err := rows.Scan(&id, &title, &description, &columnId)
		if err != nil {
			return nil, fmt.Errorf("failed scanning card rows: %w", err)
		}

		cards = append(cards, domainCard.New(id, title, description, columnId))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying cards: %w", err)
	}

	return &cards, nil
}

func (c Card) Create(ctx context.Context, title string, description string, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO cards (title, description, column_id) VALUES (?, ?, ?)", title, description, columnId)
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	return domainCard.New(int(lastInsertID), title, description, columnId), nil
}

func (c Card) Update(ctx context.Context, id int, title string, description string, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE cards SET title = ?, description = ? WHERE id = ?", title, description, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating card: %w", err)
	}

	return domainCard.New(id, title, description, columnId), nil
}

func (c Card) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting card: %w", err)
	}

	return nil
}
