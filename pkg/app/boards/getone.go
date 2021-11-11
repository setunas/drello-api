package boards

import (
	"context"
	boardDomain "drello-api/pkg/domain/board"
	cardDomain "drello-api/pkg/domain/card"
	columnDomain "drello-api/pkg/domain/column"
	"drello-api/pkg/infrastructure/repository"
)

func GetOne(ctx context.Context, boardRepo repository.Board, columnRepo repository.Column, cardRepo repository.Card, input *GetOneInput) (*GetOneOutput, error) {
	board, err := boardRepo.GetOne(ctx, input.id)
	if err != nil {
		return nil, err
	}

	columns, err := columnRepo.GetListByBoardId(ctx, board.ID())
	if err != nil {
		return nil, err
	}

	columnIds := []int{}
	for _, column := range *columns {
		columnIds = append(columnIds, column.ID())
	}
	cards, err := cardRepo.GetListByColumnIds(ctx, columnIds)
	if err != nil {
		return nil, err
	}

	return &GetOneOutput{Board: boardDomain.New(board.ID(), board.Title()), Columns: *columns, Cards: *cards}, nil
}

type GetOneInput struct {
	id int
}

func NewGetOneInput(id int) *GetOneInput {
	return &GetOneInput{id: id}
}

type GetOneOutput struct {
	Board   *boardDomain.Board
	Columns []*columnDomain.Column
	Cards   []*cardDomain.Card
}
