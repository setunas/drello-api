package updateColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func Call(ctx context.Context, columnID int, title string, position float64, boardID int, firebaseUID string) (*column.Column, error) {
	err := authorize(ctx, firebaseUID, columnID, boardID)
	if err != nil {
		return nil, err
	}

	columnDomain, err := (*repository.ColumnDS()).Update(ctx, columnID, title, position, boardID)
	if err != nil {
		return nil, err
	}

	return columnDomain, nil
}

func authorize(ctx context.Context, firebaseUID string, columnID int, boardID int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}
	if user.BoardID() != boardID {
		return fmt.Errorf("invalid board ID that you are changing to: %d, user's borad ID is: %d", boardID, user.BoardID())
	}

	column, err := (*repository.ColumnDS()).GetOneByID(ctx, columnID)
	if err != nil {
		return err
	}
	if user.BoardID() != column.BoardId() {
		return fmt.Errorf("invalid board ID that you are changing from: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	return nil
}
