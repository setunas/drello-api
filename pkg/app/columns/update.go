package columns

import (
	"context"
	"drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/repository"
)

func Update(ctx context.Context, columnRepo repository.Column, input *UpdateInput) (*UpdateOutput, error) {
	columnDomain, err := columnRepo.Update(ctx, input.id, input.title)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Column: *columnDomain}, nil
}

type UpdateInput struct {
	id    int
	title string
}

func NewUpdateInput(id int, title string) *UpdateInput {
	return &UpdateInput{id: id, title: title}
}

type UpdateOutput struct {
	Column column.Column
}
