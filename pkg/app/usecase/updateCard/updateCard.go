package updateCard

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func Call(ctx context.Context, cardID int, title string, description string, position float64, columnID int, firebaseUID string) (*card.Card, error) {
	err := authorize(ctx, firebaseUID, cardID, columnID)
	if err != nil {
		return nil, err
	}

	cardDomain, err := (*repository.CardDS()).Update(ctx, cardID, title, description, position, columnID)
	if err != nil {
		return nil, err
	}

	return cardDomain, nil
}

func authorize(ctx context.Context, firebaseUID string, cardID int, columnID int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}
	card, err := (*repository.CardDS()).GetOneByID(ctx, cardID)
	if err != nil {
		return err
	}

	oldTargetColumn, err := (*repository.ColumnDS()).GetOneByID(ctx, card.ColumnId())
	if err != nil {
		return err
	}
	if user.BoardID() != oldTargetColumn.BoardId() {
		return fmt.Errorf("invalid old target column's board ID: %d, user's borad ID is: %d", oldTargetColumn.BoardId(), user.BoardID())
	}

	newTargetColumn, err := (*repository.ColumnDS()).GetOneByID(ctx, columnID)
	if err != nil {
		return err
	}
	if user.BoardID() != newTargetColumn.BoardId() {
		return fmt.Errorf("invalid new target column's board ID: %d, user's borad ID is: %d", newTargetColumn.BoardId(), user.BoardID())
	}

	return nil
}
