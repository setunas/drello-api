package boards

import (
	"context"
	wdomain "drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/repository"
)

func GetOne(ctx context.Context, boardRepo repository.Board, input *GetOneInput) (*GetOneOutput, error) {
	board, err := boardRepo.GetOne(ctx, input.id)
	if err != nil {
		return nil, err
	}

	return &GetOneOutput{Board: wdomain.New(board.ID(), board.Title())}, nil
}

type GetOneInput struct {
	id int
}

func NewGetOneInput(id int) *GetOneInput {
	return &GetOneInput{id: id}
}

type GetOneOutput struct {
	Board *wdomain.Board
}
