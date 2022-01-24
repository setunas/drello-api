package boards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
	cardDomain "drello-api/pkg/domain/card"
	columnDomain "drello-api/pkg/domain/column"
	"fmt"
)

func GetOne(ctx context.Context, boardRepo repository.Board, columnRepo repository.Column, cardRepo repository.Card, userRepo repository.User, input *GetOneInput) (*GetOneOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != input.id {
		return nil, fmt.Errorf("not a valid request with firebase UID: %s, and board id: %d", input.firebaseUID, input.id)
	}

	board, err := boardRepo.GetOne(ctx, input.id)
	if err != nil {
		return nil, err
	}

	columns, err := columnRepo.GetListByBoardId(ctx, board.ID())
	if err != nil {
		return nil, err
	}

	columnIds := []int{}
	for _, column := range columns {
		columnIds = append(columnIds, column.ID())
	}
	cards, err := cardRepo.GetListByColumnIds(ctx, columnIds)
	if err != nil {
		return nil, err
	}

	return &GetOneOutput{Board: boardDomain.New(board.ID(), board.Title()), Columns: columns, Cards: cards}, nil
}

type GetOneInput struct {
	id          int
	firebaseUID string
}

func NewGetOneInput(id int, firebaseUID string) *GetOneInput {
	return &GetOneInput{id: id, firebaseUID: firebaseUID}
}

type GetOneOutput struct {
	Board   *boardDomain.Board
	Columns []*columnDomain.Column
	Cards   []*cardDomain.Card
}
