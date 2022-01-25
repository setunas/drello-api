package columnsDS

import (
	"context"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (c ColumnsDS) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM columns WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting column: %w", err)
	}

	return nil
}
