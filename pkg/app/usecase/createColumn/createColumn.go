package createColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func CreateColumn(ctx context.Context, title string, position float64, boardId int, firebaseUID string) (*column.Column, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != boardId {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", boardId, user.BoardID())
	}

	columnDomain, err := (*repository.ColumnDS()).Create(ctx, title, position, boardId)
	if err != nil {
		return nil, err
	}

	return columnDomain, nil
}
