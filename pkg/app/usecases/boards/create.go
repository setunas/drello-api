package boards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
)

func Create(ctx context.Context, boardRepo repository.Board, title string) (*boardDomain.Board, error) {
	board, err := boardRepo.Create(ctx, title)
	if err != nil {
		return nil, err
	}

	return board, nil
}
