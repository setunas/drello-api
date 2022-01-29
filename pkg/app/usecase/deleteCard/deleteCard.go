package deleteCard

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

	err = (*repository.CardDS()).Delete(ctx, cardID)
	if err != nil {
		return err
	}

	return nil
}

func authorize(ctx context.Context, user *user.User, cardID int) error {
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
