package repository

import (
	"context"
	"drello-api/pkg/domain/card"
)

type Card interface {
	GetOneByID(ctx context.Context, id int) (*card.Card, error)
	GetListByIDs(ctx context.Context, ids []int) ([]*card.Card, error)
	GetListByColumnIds(ctx context.Context, columnIds []int) ([]*card.Card, error)
	Create(ctx context.Context, title string, description string, position float64, columnId int) (*card.Card, error)
	Update(ctx context.Context, id int, title string, description string, position float64, columnId int) (*card.Card, error)
	UpdatePositions(ctx context.Context, data []struct {
		ID       int
		Position float64
	}) error
	Delete(ctx context.Context, id int) error
}

var cardDS Card

func SetCardDS(ds *Card) {
	cardDS = *ds
}

func CardDS() *Card {
	return &cardDS
}
