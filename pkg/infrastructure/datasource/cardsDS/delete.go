package cardsDS

import (
	"context"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (c CardsDS) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM cards WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting card: %w", err)
	}

	return nil
}
