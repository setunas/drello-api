package repository

import (
	"context"
	"drello-api/pkg/domain/card"
)

type Card interface {
	Create(ctx context.Context, title string, description string) (*card.Card, error)
	Update(ctx context.Context, id int, title string, description string) (*card.Card, error)
	Delete(ctx context.Context, id int) error
}
