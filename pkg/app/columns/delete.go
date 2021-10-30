package columns

import (
	"context"
	"drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/repository"
)

func Delete(ctx context.Context, columnRepo repository.Column, input *DeleteInput) error {
	err := columnRepo.Delete(ctx, input.id)
	if err != nil {
		return err
	}

	return nil
}

type DeleteInput struct {
	id int
}

func NewDeleteInput(id int) *DeleteInput {
	return &DeleteInput{id: id}
}

type DeleteOutput struct {
	Column column.Column
}
