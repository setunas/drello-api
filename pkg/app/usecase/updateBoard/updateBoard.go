package updateBoard

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/board"
	"drello-api/pkg/domain/user"
	"fmt"
)

func Call(ctx context.Context, boardID int, title string, user *user.User) (*board.Board, error) {
	err := authorize(ctx, user, boardID)
	if err != nil {
		return nil, err
	}

	boardDomain, err := (*repository.BoardDS()).Update(ctx, boardID, title)
	if err != nil {
		return nil, err
	}

	return boardDomain, nil
}

func authorize(ctx context.Context, user *user.User, boardID int) error {
	if user.BoardID() != boardID {
		return fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", boardID, user.BoardID())
	}

	return nil
}
