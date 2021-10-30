package datasource

import (
	"context"
	"database/sql"
	domainBoard "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

type Board struct{}

func (w Board) GetOne(ctx context.Context, id int) (*domainBoard.Board, error) {
	var title string

	db := mysql.DBPool()
	row := db.QueryRow("SELECT title FROM boards WHERE id = ?", id)

	switch err := row.Scan(&title); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainBoard.New(id, title), nil
	default:
		return nil, err
	}
}

func (w Board) Create(ctx context.Context, title string) (*domainBoard.Board, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO boards (title) VALUES (?)", title)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}

	return domainBoard.New(int(lastInsertID), title), nil
}

func (w Board) Update(ctx context.Context, id int, title string) (*domainBoard.Board, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE boards SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating board: %w", err)
	}

	return domainBoard.New(id, title), nil
}
