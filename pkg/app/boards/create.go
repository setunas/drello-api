package boards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
)

func Create(ctx context.Context, boardRepo repository.Board, input *CreateInput) (*CreateOutput, error) {
	board, err := boardRepo.Create(ctx, input.title)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Board: *board}, nil
}

type CreateInput struct {
	title string
}

func NewCreateInput(title string) *CreateInput {
	return &CreateInput{title: title}
}

type CreateOutput struct {
	Board boardDomain.Board
}
