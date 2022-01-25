package usecase

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func UpdateColumn(ctx context.Context, id int, title string, position float64, boardId int, firebaseUID string) (*column.Column, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != boardId {
		return nil, fmt.Errorf("invalid board ID that you are changing to: %d, user's borad ID is: %d", boardId, user.BoardID())
	}
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != column.BoardId() {
		return nil, fmt.Errorf("invalid board ID that you are changing from: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	columnDomain, err := (*repository.ColumnDS()).Update(ctx, id, title, position, boardId)
	if err != nil {
		return nil, err
	}

	return columnDomain, nil
}
