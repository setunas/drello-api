package boards

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/board"
	"fmt"
)

func Update(ctx context.Context, boardRepo repository.Board, userRepo repository.User, input *UpdateInput) (*UpdateOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != input.id {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", input.id, user.BoardID())
	}

	boardDomain, err := boardRepo.Update(ctx, input.id, input.title)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Board: *boardDomain}, nil
}

type UpdateInput struct {
	id          int
	title       string
	firebaseUID string
}

func NewUpdateInput(id int, title string, firebaseUID string) *UpdateInput {
	return &UpdateInput{id: id, title: title, firebaseUID: firebaseUID}
}

type UpdateOutput struct {
	Board board.Board
}
