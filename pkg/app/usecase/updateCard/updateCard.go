package updateCard

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func Call(ctx context.Context, id int, title string, description string, position float64, columnId int, firebaseUID string) (*card.Card, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}
	card, err := (*repository.CardDS()).GetOneByID(ctx, id)
	if err != nil {
		return nil, err
	}
	oldTargetColumn, err := (*repository.ColumnDS()).GetOneByID(ctx, card.ColumnId())
	if err != nil {
		return nil, err
	}
	if user.BoardID() != oldTargetColumn.BoardId() {
		return nil, fmt.Errorf("invalid old target column's board ID: %d, user's borad ID is: %d", oldTargetColumn.BoardId(), user.BoardID())
	}
	newTargetColumn, err := (*repository.ColumnDS()).GetOneByID(ctx, columnId)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != newTargetColumn.BoardId() {
		return nil, fmt.Errorf("invalid new target column's board ID: %d, user's borad ID is: %d", newTargetColumn.BoardId(), user.BoardID())
	}

	cardDomain, err := (*repository.CardDS()).Update(ctx, id, title, description, position, columnId)
	if err != nil {
		return nil, err
	}

	return cardDomain, nil
}
