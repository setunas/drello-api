package datasource

import (
	"context"
	"database/sql"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

type Column struct{}

func (c Column) GetOneByID(ctx context.Context, id int) (*domainColumn.Column, error) {
	var title string
	var position float64
	var boardID int

	db := mysql.DBPool()
	row := db.QueryRow("SELECT title, position, board_id FROM columns WHERE id = ?", id)

	switch err := row.Scan(&title, &position, &boardID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainColumn.New(id, title, position, boardID), nil
	default:
		return nil, err
	}
}

func (c Column) GetListByBoardId(ctx context.Context, boardId int) ([]*domainColumn.Column, error) {
	db := mysql.DBPool()
	rows, err := db.Query("SELECT id, title, position FROM columns WHERE board_id = ?", boardId)
	if err != nil {
		return nil, fmt.Errorf("failed querying columns: %w", err)
	}
	defer rows.Close()

	columns := []*domainColumn.Column{}

	for rows.Next() {
		var id int
		var title string
		var position float64

		err := rows.Scan(&id, &title, &position)
		if err != nil {
			return nil, fmt.Errorf("failed querying columns: %w", err)
		}

		columns = append(columns, domainColumn.New(id, title, position, boardId))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying columns: %w", err)
	}

	return columns, nil
}

func (c Column) Create(ctx context.Context, title string, position float64, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO columns (title, position, board_id) VALUES (?, ?, ?)", title, position, boardId)
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating column: %w", err)
	}

	return domainColumn.New(int(lastInsertID), title, position, boardId), nil
}

func (c Column) Update(ctx context.Context, id int, title string, position float64, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE columns SET title = ?, position = ? WHERE id = ?", title, position, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating column: %w", err)
	}

	return domainColumn.New(id, title, position, boardId), nil
}

func (c Column) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM columns WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting column: %w", err)
	}

	return nil
}
