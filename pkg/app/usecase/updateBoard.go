package usecase

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/board"
	"fmt"
)

func UpdateBoard(ctx context.Context, id int, title string, firebaseUID string) (*board.Board, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != id {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", id, user.BoardID())
	}

	boardDomain, err := (*repository.BoardDS()).Update(ctx, id, title)
	if err != nil {
		return nil, err
	}

	return boardDomain, nil
}
