package columnsDS

import (
	"context"
	"database/sql"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c ColumnsDS) GetOneByID(ctx context.Context, id int) (*domainColumn.Column, error) {
	var title string
	var position float64
	var boardID int

	db := mysql.DBPool()
	query := "SELECT title, position, board_id FROM columns WHERE id = ?"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("id", id).Write()
	row := db.QueryRow(query, id)

	switch err := row.Scan(&title, &position, &boardID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainColumn.New(id, title, position, boardID), nil
	default:
		return nil, err
	}
}
