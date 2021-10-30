package boards

import (
	"context"
	"drello-api/pkg/domain/board"
	"drello-api/pkg/infrastructure/repository"
)

func Update(ctx context.Context, boardRepo repository.Board, input *UpdateInput) (*UpdateOutput, error) {
	boardDomain, err := boardRepo.Update(ctx, input.id, input.title)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Board: *boardDomain}, nil
}

type UpdateInput struct {
	id    int
	title string
}

func NewUpdateInput(id int, title string) *UpdateInput {
	return &UpdateInput{id: id, title: title}
}

type UpdateOutput struct {
	Board board.Board
}
