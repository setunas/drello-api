package createCard

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"drello-api/pkg/domain/user"
	"fmt"
)

func Call(ctx context.Context, title string, description string, position float64, columnId int, user *user.User) (*card.Card, error) {
	err := authorize(ctx, user, columnId)
	if err != nil {
		return nil, err
	}

	cardDomain, err := (*repository.CardDS()).Create(ctx, title, description, position, columnId)
	if err != nil {
		return nil, err
	}

	return cardDomain, nil
}

func authorize(ctx context.Context, user *user.User, columnId int) error {
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, columnId)
	if err != nil {
		return err
	}

	if user.BoardID() != column.BoardId() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	return nil
}
