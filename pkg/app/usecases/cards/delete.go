package cards

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func Delete(ctx context.Context, columnRepo repository.Column, cardRepo repository.Card, userRepo repository.User, input *DeleteInput) error {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return err
	}
	card, err := cardRepo.GetOneByID(ctx, input.id)
	if err != nil {
		return err
	}
	column, err := columnRepo.GetOneByID(ctx, card.ColumnId())
	if err != nil {
		return err
	}
	if user.BoardID() != column.BoardId() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	err = cardRepo.Delete(ctx, input.id)
	if err != nil {
		return err
	}

	return nil
}

type DeleteInput struct {
	id          int
	firebaseUID string
}

func NewDeleteInput(id int, firebaseUID string) *DeleteInput {
	return &DeleteInput{id: id, firebaseUID: firebaseUID}
}

type DeleteOutput struct {
	Card card.Card
}
