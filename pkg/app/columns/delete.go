package columns

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func Delete(ctx context.Context, boardRepo repository.Board, columnRepo repository.Column, userRepo repository.User, input *DeleteInput) error {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return err
	}

	board, err := boardRepo.GetOne(ctx, input.id)
	if err != nil {
		return err
	}
	if user.BoardID() != board.ID() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", board.ID(), user.BoardID())
	}

	err = columnRepo.Delete(ctx, input.id)
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
	Column column.Column
}
