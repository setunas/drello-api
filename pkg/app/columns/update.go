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
		return nil, fmt.Errorf("invalid board ID that you are changing to: %d, user's borad ID is: %d", input.boardId, user.BoardID())
	}
	column, err := columnRepo.GetOneByID(ctx, input.id)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != column.BoardId() {
		return nil, fmt.Errorf("invalid board ID that you are changing from: %d, user's borad ID is: %d", column.BoardId(), user.BoardID())
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
	position    float64
	boardId     int
	firebaseUID string
}

func NewUpdateInput(id int, title string, position float64, boardId int, firebaseUID string) *UpdateInput {
	return &UpdateInput{id: id, title: title, position: position, boardId: boardId, firebaseUID: firebaseUID}
}

type UpdateOutput struct {
	Column column.Column
}
