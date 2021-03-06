package repository

import (
	"context"
	"drello-api/pkg/domain/column"
)

type Column interface {
	GetOneByID(ctx context.Context, columnID int) (*column.Column, error)
	GetListByBoardId(ctx context.Context, boardId int) ([]*column.Column, error)
	Create(ctx context.Context, title string, position float64, boardId int) (*column.Column, error)
	Update(ctx context.Context, id int, title string, position float64, boardId int) (*column.Column, error)
	Delete(ctx context.Context, id int) error
}

var columnDS *Column

func SetColumnDS(ds Column) {
	columnDS = &ds
}

func ColumnDS() *Column {
	return columnDS
}
