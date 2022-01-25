package getBoardWithColumnsAndCards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
	cardDomain "drello-api/pkg/domain/card"
	columnDomain "drello-api/pkg/domain/column"
	"fmt"
)

func Call(ctx context.Context, id int, firebaseUID string) (*boardDomain.Board, []*columnDomain.Column, []*cardDomain.Card, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, nil, nil, err
	}
	if user.BoardID() != id {
		return nil, nil, nil, fmt.Errorf("not a valid request with firebase UID: %s, and board id: %d", firebaseUID, id)
	}

	board, err := (*repository.BoardDS()).GetOne(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	columns, err := (*repository.ColumnDS()).GetListByBoardId(ctx, board.ID())
	if err != nil {
		return nil, nil, nil, err
	}

	columnIds := []int{}
	for _, column := range columns {
		columnIds = append(columnIds, column.ID())
	}
	cards, err := (*repository.CardDS()).GetListByColumnIds(ctx, columnIds)
	if err != nil {
		return nil, nil, nil, err
	}

	return boardDomain.New(board.ID(), board.Title()), columns, cards, nil
}
