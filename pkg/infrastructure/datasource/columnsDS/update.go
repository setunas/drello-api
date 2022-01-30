package columnsDS

import (
	"context"
	domainColumn "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/log"
	"fmt"
)

func (c ColumnsDS) Update(ctx context.Context, id int, title string, position float64, boardId int) (*domainColumn.Column, error) {
	db := mysql.DBPool()
	query := "UPDATE columns SET title = ?, position = ? WHERE id = ?"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("title", title).Add("position", position).Write()
	_, err := db.Exec(query, title, position, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating column: %w", err)
	}

	return domainColumn.New(id, title, position, boardId), nil
}
