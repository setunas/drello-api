package columns

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func Update(ctx context.Context, columnRepo repository.Column, userRepo repository.User, input *UpdateInput) (*UpdateOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != input.boardId {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", input.boardId, user.BoardID())
	}

	columnDomain, err := columnRepo.Update(ctx, input.id, input.title, input.boardId)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Column: *columnDomain}, nil
}

type UpdateInput struct {
	id          int
	title       string
	boardId     int
	firebaseUID string
}

func NewUpdateInput(id int, title string, boardId int, firebaseUID string) *UpdateInput {
	return &UpdateInput{id: id, title: title, boardId: boardId, firebaseUID: firebaseUID}
}

type UpdateOutput struct {
	Column column.Column
}
