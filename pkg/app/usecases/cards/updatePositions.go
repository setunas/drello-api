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

	columns, err := columnRepo.GetListByBoardId(ctx, user.BoardID())
	if err != nil {
		return nil, err
	}

	columnMap := make(map[int]bool)
	for _, column := range *columns {
		columnMap[column.ID()] = true
	}

	cardIDs := make([]int, len(input.cards))
	for _, card := range input.cards {
		cardIDs = append(cardIDs, card.id)
	}
	cards, err := cardRepo.GetListByIDs(ctx, cardIDs)
	if err != nil {
		return nil, err
	}

	for _, card := range *cards {
		if columnMap[card.ColumnId()] != true {
			return nil, fmt.Errorf("Invalid columnID: %d", card.ColumnId())
		}
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
