package cards

import (
	"context"
	"drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/repository"
)

func Update(ctx context.Context, cardRepo repository.Card, input *UpdateInput) (*UpdateOutput, error) {
	cardDomain, err := cardRepo.Update(ctx, input.id, input.title, input.description)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Card: *cardDomain}, nil
}

type UpdateInput struct {
	id          int
	title       string
	description string
}

func NewUpdateInput(id int, title string, description string) *UpdateInput {
	return &UpdateInput{id: id, title: title, description: description}
}

type UpdateOutput struct {
	Card card.Card
}
