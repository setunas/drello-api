package columns

import (
	"context"
	"drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/repository"
)

func Update(ctx context.Context, columnRepo repository.Column, input *UpdateInput) (*UpdateOutput, error) {
	columnDomain, err := columnRepo.Update(ctx, input.id, input.title, input.boardId)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Column: *columnDomain}, nil
}

type UpdateInput struct {
	id      int
	title   string
	boardId int
}

func NewUpdateInput(id int, title string, boardId int) *UpdateInput {
	return &UpdateInput{id: id, title: title, boardId: boardId}
}

type UpdateOutput struct {
	Column column.Column
}
