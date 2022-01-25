package signup

import (
	"context"
	"drello-api/pkg/app/repository"
	columnDM "drello-api/pkg/domain/column"
	userDM "drello-api/pkg/domain/user"
)

func Call(ctx context.Context, username string, firebaseUID string, title string) (*userDM.User, error) {
	board, err := (*repository.BoardDS()).Create(ctx, title)
	if err != nil {
		return nil, err
	}

	user, err := (*repository.UserDS()).Create(ctx, username, board.ID(), firebaseUID)
	if err != nil {
		return nil, err
	}

	_, err = (*repository.ColumnDS()).Create(ctx, "Todo", columnDM.InitialPositionGap()*1, board.ID())
	if err != nil {
		return nil, err
	}
	_, err = (*repository.ColumnDS()).Create(ctx, "Doing", columnDM.InitialPositionGap()*2, board.ID())
	if err != nil {
		return nil, err
	}
	_, err = (*repository.ColumnDS()).Create(ctx, "Done", columnDM.InitialPositionGap()*3, board.ID())
	if err != nil {
		return nil, err
	}

	return user, nil
}
