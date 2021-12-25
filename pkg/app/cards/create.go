package cards

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func Create(ctx context.Context, columnRepo repository.Column, cardRepo repository.Card, userRepo repository.User, input *CreateInput) (*CreateOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}
	column, err := columnRepo.GetOneById(ctx, input.columnId)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != column.BoardId() {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

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
	firebaseUID string
}

func NewCreateInput(title string, description string, columnId int, firebaseUID string) *CreateInput {
	return &CreateInput{title: title, description: description, columnId: columnId, firebaseUID: firebaseUID}
}

type CreateOutput struct {
	Card card.Card
}
