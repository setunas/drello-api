package cards

import (
	"context"
	"drello-api/pkg/app/repository"
	"fmt"
)

func UpdatePositions(ctx context.Context, input *UpdatePositionsInput) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return err
	}

	columns, err := (*repository.ColumnDS()).GetListByBoardId(ctx, user.BoardID())
	if err != nil {
		return err
	}

	columnMap := make(map[int]bool)
	for _, column := range columns {
		columnMap[column.ID()] = true
	}

	cardIDs := make([]int, len(input.cards))
	for _, card := range input.cards {
		cardIDs = append(cardIDs, card.id)
	}
	cards, err := (*repository.CardDS()).GetListByIDs(ctx, cardIDs)
	if err != nil {
		return err
	}

	for _, card := range cards {
		if columnMap[card.ColumnId()] != true {
			return fmt.Errorf("invalid columnID: %d", card.ColumnId())
		}
	}

	data := make([]struct {
		ID       int
		Position float64
	}, len(input.cards))
	for i, c := range input.cards {
		data[i].ID = c.id
		data[i].Position = c.position
	}

	err = (*repository.CardDS()).UpdatePositions(ctx, data)
	if err != nil {
		return err
	}

	return nil
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
