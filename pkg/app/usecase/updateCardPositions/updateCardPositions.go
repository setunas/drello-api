package updateCardPositions

import (
	"context"
	"drello-api/pkg/app/repository"
	"fmt"
)

func Call(ctx context.Context, inputCards []Card, firebaseUID string) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
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

	cardIDs := make([]int, len(inputCards))
	for _, card := range inputCards {
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
	}, len(inputCards))
	for i, c := range inputCards {
		data[i].ID = c.id
		data[i].Position = c.position
	}

	err = (*repository.CardDS()).UpdatePositions(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

type Card struct {
	id       int
	position float64
}

func NewCard(id int, position float64) *Card {
	return &Card{id: id, position: position}
}
