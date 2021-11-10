package cards

import (
	"context"
	"drello-api/pkg/domain/card"
	"drello-api/pkg/infrastructure/repository"
)

func Create(ctx context.Context, cardRepo repository.Card, input *CreateInput) (*CreateOutput, error) {
	cardDomain, err := cardRepo.Create(ctx, input.title, input.description, input.columnId)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Card: *cardDomain}, nil
}

type CreateInput struct {
	title       string
	description string
	columnId    int
}

func NewCreateInput(title string, description string, columnId int) *CreateInput {
	return &CreateInput{title: title, description: description, columnId: columnId}
}

type CreateOutput struct {
	Card card.Card
}
