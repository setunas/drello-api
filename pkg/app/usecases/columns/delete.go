package columns

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func Delete(ctx context.Context, input *DeleteInput) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return err
	}
	column, err := (*repository.ColumnDS()).GetOneByID(ctx, input.id)
	if err != nil {
		return err
	}
	board, err := (*repository.BoardDS()).GetOne(ctx, column.BoardId())
	if err != nil {
		return err
	}
	if user.BoardID() != board.ID() {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", board.ID(), user.BoardID())
	}

	err = (*repository.ColumnDS()).Delete(ctx, input.id)
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
