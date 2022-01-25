package createCard

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func Call(ctx context.Context, title string, description string, position float64, columnId int, firebaseUID string) (*card.Card, error) {
	err := authorize(ctx, firebaseUID, columnId)
	if err != nil {
		return nil, err
	}

	cardDomain, err := (*repository.CardDS()).Create(ctx, title, description, position, columnId)
	if err != nil {
		return nil, err
	}

	return cardDomain, nil
}

func authorize(ctx context.Context, firebaseUID string, columnId int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, columnId)
	if err != nil {
		return err
	}

	if user.BoardID() != column.BoardId() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	return nil
}
