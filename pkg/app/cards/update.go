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
	card, err := cardRepo.GetOneByID(ctx, input.id)
	if err != nil {
		return nil, err
	}
	oldTargetColumn, err := columnRepo.GetOneByID(ctx, card.ColumnId())
	if err != nil {
		return nil, err
	}
	if user.BoardID() != oldTargetColumn.BoardId() {
		return nil, fmt.Errorf("invalid old target column's board ID: %d, user's borad ID is: %d", oldTargetColumn.BoardId(), user.BoardID())
	}
	newTargetColumn, err := columnRepo.GetOneByID(ctx, input.columnId)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != newTargetColumn.BoardId() {
		return nil, fmt.Errorf("invalid new target column's board ID: %d, user's borad ID is: %d", newTargetColumn.BoardId(), user.BoardID())
	}

	cardDomain, err := cardRepo.Update(ctx, input.id, input.title, input.description, input.position, input.columnId)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Card: *cardDomain}, nil
}

type UpdateInput struct {
	id          int
	title       string
	description string
	position    float64
	columnId    int
	firebaseUID string
}

func NewUpdateInput(id int, title string, description string, position float64, columnId int, firebaseUID string) *UpdateInput {
	return &UpdateInput{id: id, title: title, description: description, position: position, columnId: columnId, firebaseUID: firebaseUID}
}

type UpdateOutput struct {
	Card card.Card
}
