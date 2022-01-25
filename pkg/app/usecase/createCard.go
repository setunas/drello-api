package usecase

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func CreateCard(ctx context.Context, title string, description string, position float64, columnId int, firebaseUID string) (*card.Card, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, columnId)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != column.BoardId() {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	cardDomain, err := (*repository.CardDS()).Create(ctx, title, description, position, columnId)
	if err != nil {
		return nil, err
	}

	return cardDomain, nil
}
