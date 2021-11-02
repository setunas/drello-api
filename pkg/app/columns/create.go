package columns

import (
	"context"
	"drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/repository"
)

func Create(ctx context.Context, columnRepo repository.Column, input *CreateInput) (*CreateOutput, error) {
	columnDomain, err := columnRepo.Create(ctx, input.title)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Column: *columnDomain}, nil
}

type CreateInput struct {
	title string
}

func NewCreateInput(title string) *CreateInput {
	return &CreateInput{title: title}
}

type CreateOutput struct {
	Column column.Column
}
