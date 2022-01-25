package createColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func Call(ctx context.Context, title string, position float64, boardID int, firebaseUID string) (*column.Column, error) {
	err := authorize(ctx, firebaseUID, boardID)
	if err != nil {
		return nil, err
	}

	columnDomain, err := (*repository.ColumnDS()).Create(ctx, title, position, boardID)
	if err != nil {
		return nil, err
	}

	return columnDomain, nil
}

func authorize(ctx context.Context, firebaseUID string, boardID int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}

	if user.BoardID() != boardID {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", boardID, user.BoardID())
	}

	return nil
}
