package columns

import (
	"context"
	"drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/repository"
)

func Create(ctx context.Context, columnRepo repository.Column, input *CreateInput) (*CreateOutput, error) {
	columnDomain, err := columnRepo.Create(ctx, input.title, input.boardId)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Column: *columnDomain}, nil
}

type CreateInput struct {
	title   string
	boardId int
}

func NewCreateInput(title string, boardId int) *CreateInput {
	return &CreateInput{title: title, boardId: boardId}
}

type CreateOutput struct {
	Column column.Column
}
