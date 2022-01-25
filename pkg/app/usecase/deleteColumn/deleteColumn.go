package deleteColumn

import (
	"context"
	"drello-api/pkg/app/repository"
	"fmt"
)

func DeleteColumn(ctx context.Context, id int, firebaseUID string) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, id)
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

	err = (*repository.ColumnDS()).Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
