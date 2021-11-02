package repository

import (
	"context"
	"drello-api/pkg/domain/column"
)

type Column interface {
	Create(ctx context.Context, title string) (*column.Column, error)
	Update(ctx context.Context, id int, title string) (*column.Column, error)
	Delete(ctx context.Context, id int) error
}
