package deleteColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"fmt"
)

func Call(ctx context.Context, cardID int, firebaseUID string) error {
	err := authorize(ctx, firebaseUID, cardID)
	if err != nil {
		return err
	}

	err = (*repository.ColumnDS()).Delete(ctx, cardID)
	if err != nil {
		return err
	}

	return nil
}

func authorize(ctx context.Context, firebaseUID string, cardID int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}
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
