package columnsDS

import (
	"context"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c ColumnsDS) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	query := "DELETE FROM columns WHERE id = ?"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("id", id).Write()
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed deleting column: %w", err)
	}

	return nil
}
