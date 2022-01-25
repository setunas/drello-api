package cardsDS

import (
	"context"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
	"strings"
)

func (c CardsDS) UpdatePositions(ctx context.Context, data []struct {
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
