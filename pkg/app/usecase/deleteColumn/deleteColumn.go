package deleteColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/user"
	"fmt"
)

func Call(ctx context.Context, cardID int, user *user.User) error {
	err := authorize(ctx, user, cardID)
	if err != nil {
		return err
	}

	err = (*repository.ColumnDS()).Delete(ctx, cardID)
	if err != nil {
		return err
	}

	return nil
}

func authorize(ctx context.Context, user *user.User, cardID int) error {
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, cardID)
	if err != nil {
		return err
	}
	board, err := (*repository.BoardDS()).GetOne(ctx, column.BoardId())
	if err != nil {
		return err
	}

	if user.BoardID() != board.ID() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", board.ID(), user.BoardID())
	}

	return nil
}
