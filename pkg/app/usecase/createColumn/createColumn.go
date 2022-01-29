package createColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"drello-api/pkg/domain/user"
	"fmt"
)

func Call(ctx context.Context, title string, position float64, boardID int, user *user.User) (*column.Column, error) {
	err := authorize(ctx, user, boardID)
	if err != nil {
		return nil, err
	}

	columnDomain, err := (*repository.ColumnDS()).Create(ctx, title, position, boardID)
	if err != nil {
		return nil, err
	}

	return columnDomain, nil
}

func authorize(ctx context.Context, user *user.User, boardID int) error {
	if user.BoardID() != boardID {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", boardID, user.BoardID())
	}

	return nil
}
