package datasource

import (
	"context"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

type Column struct{}

func (c Column) Create(ctx context.Context, title string) (*domainColumn.Column, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO columns (title) VALUES (?)", title)
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	return domainColumn.New(int(lastInsertID), title), nil
}

func (c Column) Update(ctx context.Context, id int, title string) (*domainColumn.Column, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE columns SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating column: %w", err)
	}

	return domainColumn.New(id, title), nil
}

func (c Column) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM columns WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting column: %w", err)
	}

	return nil
}
