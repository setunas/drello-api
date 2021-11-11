package repository

import (
	"context"
	"drello-api/pkg/domain/card"
)

type Card interface {
	GetListByColumnIds(ctx context.Context, columnId []int) (*[]*card.Card, error)
	Create(ctx context.Context, title string, description string, columnId int) (*card.Card, error)
	Update(ctx context.Context, id int, title string, description string, columnId int) (*card.Card, error)
	Delete(ctx context.Context, id int) error
}
