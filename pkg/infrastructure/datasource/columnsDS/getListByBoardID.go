package columnsDS

import (
	"context"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c ColumnsDS) GetListByBoardId(ctx context.Context, boardId int) ([]*domainColumn.Column, error) {
	db := mysql.DBPool()
	query := "SELECT id, title, position FROM columns WHERE board_id = ?"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("boardId", boardId).Write()
	rows, err := db.Query(query, boardId)
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
