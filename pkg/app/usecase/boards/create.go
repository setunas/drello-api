package boards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
)

func Create(ctx context.Context, title string) (*boardDomain.Board, error) {
	board, err := (*repository.BoardDS()).Create(ctx, title)
	if err != nil {
		return nil, err
	}

	return board, nil
}
