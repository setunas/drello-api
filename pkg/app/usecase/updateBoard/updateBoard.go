package updateBoard

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/board"
	"fmt"
)

func Call(ctx context.Context, boardID int, title string, firebaseUID string) (*board.Board, error) {
	err := authorize(ctx, firebaseUID, boardID)
	if err != nil {
		return nil, err
	}

	boardDomain, err := (*repository.BoardDS()).Update(ctx, boardID, title)
	if err != nil {
		return nil, err
	}

	return boardDomain, nil
}

func authorize(ctx context.Context, firebaseUID string, boardID int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}

	if user.BoardID() != boardID {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", boardID, user.BoardID())
	}

	return nil
}
