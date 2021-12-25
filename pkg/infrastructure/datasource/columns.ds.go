package datasource

import (
	"context"
	"database/sql"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

type Column struct{}

func (c Column) GetOne(ctx context.Context, id int) (*domainColumn.Column, error) {
	var title string
	var boardID int

	db := mysql.DBPool()
	row := db.QueryRow("SELECT title, board_id FROM column WHERE id = ?", id)

	switch err := row.Scan(&title, &boardID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainColumn.New(id, title, boardID), nil
	default:
		return nil, err
	}
}

func (c Column) GetListByBoardId(ctx context.Context, boardId int) (*[]*domainColumn.Column, error) {
	db := mysql.DBPool()
	rows, err := db.Query("SELECT id, title FROM columns WHERE board_id = ?", boardId)
	if err != nil {
		return nil, fmt.Errorf("failed querying columns: %w", err)
	}
	defer rows.Close()

	columns := []*domainColumn.Column{}

	for rows.Next() {
		var id int
		var title string

		err := rows.Scan(&id, &title)
		if err != nil {
			return nil, fmt.Errorf("failed querying columns: %w", err)
		}

		columns = append(columns, domainColumn.New(id, title, boardId))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying columns: %w", err)
	}

	return &columns, nil
}

func (c Column) Create(ctx context.Context, title string, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO columns (title, board_id) VALUES (?, ?)", title, boardId)
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	return domainColumn.New(int(lastInsertID), title, boardId), nil
}

func (c Column) Update(ctx context.Context, id int, title string, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE columns SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating column: %w", err)
	}

	return domainColumn.New(id, title, boardId), nil
}

func (c Column) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM columns WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting column: %w", err)
	}

	return nil
}
