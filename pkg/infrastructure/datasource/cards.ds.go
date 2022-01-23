package datasource

import (
	"context"
	"database/sql"
	domainCard "drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
	"strings"
)

type Card struct{}

func (c Card) GetOneByID(ctx context.Context, id int) (*domainCard.Card, error) {
	var title string
	var description string
	var position float64
	var columnID int

	db := mysql.DBPool()
	row := db.QueryRow("SELECT title, description, position, column_id FROM cards WHERE id = ?", id)

	switch err := row.Scan(&title, &description, &position, &columnID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainCard.New(id, title, description, position, columnID), nil
	default:
		return nil, err
	}
}

func (c Card) GetListByIDs(ctx context.Context, ids []int) (*[]*domainCard.Card, error) {
	db := mysql.DBPool()

	cards := []*domainCard.Card{}

	if len(ids) == 0 {
		return &cards, nil
	}

	args := make([]interface{}, len(ids))
	for i := range ids {
		args[i] = ids[i]
	}

	sql := "SELECT id, title, description, position, column_id FROM cards WHERE id IN (?" + strings.Repeat(",?", len(ids)-1) + ")"

	rows, err := db.Query(sql, args...)
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

	return &cards, nil
}

func (c Card) GetListByColumnIds(ctx context.Context, columnIds []int) (*[]*domainCard.Card, error) {
	db := mysql.DBPool()

	cards := []*domainCard.Card{}

	if len(columnIds) == 0 {
		return &cards, nil
	}

	args := make([]interface{}, len(columnIds))
	for i := range columnIds {
		args[i] = columnIds[i]
	}

	sql := "SELECT id, title, description, position, column_id FROM cards WHERE column_id IN (?" + strings.Repeat(",?", len(columnIds)-1) + ")"

	rows, err := db.Query(sql, args...)
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

	return &cards, nil
}

func (c Card) Create(ctx context.Context, title string, description string, position float64, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO cards (title, description, position, column_id) VALUES (?, ?, ?, ?)", title, description, position, columnId)
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating card: %w", err)
	}

	return domainCard.New(int(lastInsertID), title, description, position, columnId), nil
}

func (c Card) Update(ctx context.Context, id int, title string, description string, position float64, columnId int) (*domainCard.Card, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE cards SET title = ?, description = ?, position = ?, column_id = ? WHERE id = ?", title, description, position, columnId, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating card: %w", err)
	}

	return domainCard.New(id, title, description, position, columnId), nil
}

func (c Card) UpdatePositions(ctx context.Context, data []struct {
	ID       int
	Position float64
}) error {
	db := mysql.DBPool()

	if len(data) == 0 {
		return nil
	}

	ids := make([]interface{}, len(data))
	placeholders := make([]interface{}, len(data)*2)
	for i := range data {
		ids[i] = data[i].ID
		placeholders[i*2] = data[i].ID
		placeholders[i*2+1] = data[i].Position
	}
	placeholders = append(placeholders, ids...)

	sql := `
	UPDATE cards
	SET position = 
	CASE id 
	` +
		strings.Repeat("WHEN ? THEN ? ", len(data)) +
		`ELSE position END
	WHERE id IN (?` + strings.Repeat(",?", len(ids)-1) + ")"
	fmt.Println(sql)

	_, err := db.Exec(sql, placeholders...)
	if err != nil {
		return fmt.Errorf("failed updating card positions: %w", err)
	}

	return nil
}

func (c Card) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting card: %w", err)
	}

	return nil
}
