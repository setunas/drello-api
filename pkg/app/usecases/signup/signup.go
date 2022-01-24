package signup

import (
	"context"
	"drello-api/pkg/app/repository"
	columnDM "drello-api/pkg/domain/column"
	userDM "drello-api/pkg/domain/user"
)

func Signup(ctx context.Context, input *SignupInput) (*SignupOutput, error) {
	board, err := (*repository.BoardDS()).Create(ctx, input.title)
	if err != nil {
		return nil, err
	}

	user, err := (*repository.UserDS()).Create(ctx, input.username, board.ID(), input.firebaseUID)
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

	return &SignupOutput{User: *user}, nil
}

type SignupInput struct {
	title       string
	username    string
	firebaseUID string
}

func NewSignupInput(username string, firebaseUID string, title string) *SignupInput {
	return &SignupInput{username: username, firebaseUID: firebaseUID, title: title}
}

type SignupOutput struct {
	User userDM.User
}
