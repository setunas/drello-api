package cards

import (
	"context"
	"drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/repository"
)

func Delete(ctx context.Context, cardRepo repository.Card, input *DeleteInput) error {
	err := cardRepo.Delete(ctx, input.id)
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
	Card card.Card
}
