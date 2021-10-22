package repository

import (
	"context"
	"drello-api/pkg/domain/board"
)

type Board interface {
	GetOne(ctx context.Context, id int) (*board.Board, error)
	Create(ctx context.Context, title string) (*board.Board, error)
	Update(ctx context.Context, id int, title string) (*board.Board, error)
}
