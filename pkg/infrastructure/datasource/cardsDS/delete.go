package cardsDS

import (
	"context"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c CardsDS) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	query := "DELETE FROM cards WHERE id = ?"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("id", id).Write()
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed deleting card: %w", err)
	}

	return nil
}
