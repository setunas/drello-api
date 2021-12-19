package users

import (
	"context"
	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/repository"
)

func Create(ctx context.Context, userRepo repository.User, input *CreateInput) (*CreateOutput, error) {
	user, err := userRepo.Create(ctx, input.username, input.boardID, input.firebaseUID)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{User: *user}, nil
}

type CreateInput struct {
	username    string
	boardID     int
	firebaseUID string
}

func NewCreateInput(username string, boardID int, firebaseUID string) *CreateInput {
	return &CreateInput{username: username, boardID: boardID, firebaseUID: firebaseUID}
}

type CreateOutput struct {
	User userDM.User
}
