package deleteCard

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

	err = (*repository.CardDS()).Delete(ctx, cardID)
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
	card, err := (*repository.CardDS()).GetOneByID(ctx, cardID)
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

	return nil
}
