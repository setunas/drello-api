package cards

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func Update(ctx context.Context, columnRepo repository.Column, cardRepo repository.Card, userRepo repository.User, input *UpdateInput) (*UpdateOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}
	column, err := columnRepo.GetOneByID(ctx, input.columnId)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != column.BoardId() {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
	}

	cardDomain, err := cardRepo.Update(ctx, input.id, input.title, input.description, input.columnId)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Card: *cardDomain}, nil
}

type UpdateInput struct {
	id          int
	title       string
	description string
	columnId    int
	firebaseUID string
}

func NewUpdateInput(id int, title string, description string, columnId int, firebaseUID string) *UpdateInput {
	return &UpdateInput{id: id, title: title, description: description, columnId: columnId, firebaseUID: firebaseUID}
}

type UpdateOutput struct {
	Card card.Card
}
