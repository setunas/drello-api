package cards

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/card"
	"fmt"
)

func UpdatePositions(ctx context.Context, columnRepo repository.Column, cardRepo repository.Card, userRepo repository.User, input *UpdatePositionsInput) (*UpdateOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}

	///////////////////
	// Authorization //
	///////////////////
	// TODO: Get all column records of the board
	columnRepo.GetListByBoardId()
	// TODO: Set all column IDs to a map
	// TODO: Create a new repo and a datasource that return all card records of given card IDs
	// TODO: Get all card records of the provided card ID
	// TODO: Check if all card's column IDs match column IDs in the map for authorization

	oldTargetColumn, err := columnRepo.GetOneByID(ctx, input.columnID)
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

type CardInput struct {
	id       int
	position float64
}

func NewCardInput(id int, position float64) *CardInput {
	return &CardInput{id: id, position: position}
}

type UpdatePositionsInput struct {
	cards       []CardInput
	firebaseUID string
}

func NewUpdatePositionsInput(cards []CardInput, firebaseUID string) *UpdatePositionsInput {
	return &UpdatePositionsInput{cards: cards, firebaseUID: firebaseUID}
}

type UpdatePositionsOutput struct {
	Card card.Card
}
