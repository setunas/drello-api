package deleteCard

import (
	"context"
	"drello-api/pkg/app/repository"
	"fmt"
)

func Call(ctx context.Context, id int, firebaseUID string) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}
	card, err := (*repository.CardDS()).GetOneByID(ctx, id)
	if err != nil {
		return err
	}
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, card.ColumnId())
	if err != nil {
		return err
	}
	if user.BoardID() != column.BoardId() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	err = (*repository.CardDS()).Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
