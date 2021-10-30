package cards

import (
	"context"
	"drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/repository"
)

func Create(ctx context.Context, cardRepo repository.Card, input *CreateInput) (*CreateOutput, error) {
	cardDomain, err := cardRepo.Create(ctx, input.title, input.description)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Card: *cardDomain}, nil
}

type CreateInput struct {
	title       string
	description string
}

func NewCreateInput(title string, description string) *CreateInput {
	return &CreateInput{title: title, description: description}
}

type CreateOutput struct {
	Card card.Card
}
