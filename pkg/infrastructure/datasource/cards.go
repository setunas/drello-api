package datasource

import (
	"context"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

type Card struct{}

func (c Card) Create(ctx context.Context, title string, description string) (*domainCard.Card, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO cards (title, description) VALUES (?, ?)", title, description)
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	return domainCard.New(int(lastInsertID), title, description), nil
}

func (c Card) Update(ctx context.Context, id int, title string, description string) (*domainCard.Card, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE cards SET title = ?, description = ? WHERE id = ?", title, description, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating card: %w", err)
	}

	return domainCard.New(id, title, description), nil
}

func (c Card) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting card: %w", err)
	}

	return nil
}
