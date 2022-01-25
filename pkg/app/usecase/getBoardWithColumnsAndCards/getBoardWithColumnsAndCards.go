package getBoardWithColumnsAndCards

import (
	"context"
	"drello-api/pkg/app/repository"
	boardDomain "drello-api/pkg/domain/board"
	cardDomain "drello-api/pkg/domain/card"
	columnDomain "drello-api/pkg/domain/column"
	"fmt"
)

func Call(ctx context.Context, boardID int, firebaseUID string) (*boardDomain.Board, []*columnDomain.Column, []*cardDomain.Card, error) {
	err := authorize(ctx, firebaseUID, boardID)
	if err != nil {
		return nil, nil, nil, err
	}

	board, err := (*repository.BoardDS()).GetOne(ctx, boardID)
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

func authorize(ctx context.Context, firebaseUID string, boardID int) error {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return err
	}

	if user.BoardID() != boardID {
		return fmt.Errorf("not a valid request with firebase UID: %s, and board id: %d", firebaseUID, boardID)
	}

	return nil
}
