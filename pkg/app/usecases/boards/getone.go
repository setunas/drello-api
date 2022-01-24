package boards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
	cardDomain "drello-api/pkg/domain/card"
	columnDomain "drello-api/pkg/domain/column"
	"fmt"
)

func GetOne(ctx context.Context, boardRepo repository.Board, columnRepo repository.Column, cardRepo repository.Card, userRepo repository.User, id int, firebaseUID string) (*boardDomain.Board, []*columnDomain.Column, []*cardDomain.Card, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, nil, nil, err
	}
	if user.BoardID() != id {
		return nil, nil, nil, fmt.Errorf("not a valid request with firebase UID: %s, and board id: %d", firebaseUID, id)
	}

	board, err := boardRepo.GetOne(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	columns, err := columnRepo.GetListByBoardId(ctx, board.ID())
	if err != nil {
		return nil, nil, nil, err
	}

	columnIds := []int{}
	for _, column := range columns {
		columnIds = append(columnIds, column.ID())
	}
	cards, err := cardRepo.GetListByColumnIds(ctx, columnIds)
	if err != nil {
		return nil, nil, nil, err
	}

	return boardDomain.New(board.ID(), board.Title()), columns, cards, nil
}
