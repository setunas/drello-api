package repository

import (
	"context"
	"drello-api/pkg/domain/column"
)

type Column interface {
	GetOneByID(ctx context.Context, columnID int) (*column.Column, error)
	GetListByBoardId(ctx context.Context, boardId int) (*[]*column.Column, error)
	Create(ctx context.Context, title string, boardId int) (*column.Column, error)
	Update(ctx context.Context, id int, title string, boardId int) (*column.Column, error)
	Delete(ctx context.Context, id int) error
}
