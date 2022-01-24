package columns

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/column"
	"fmt"
)

func Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}
	if user.BoardID() != input.boardId {
		return nil, fmt.Errorf("invalid board ID: %d, user's borad ID is: %d", input.boardId, user.BoardID())
	}

	columnDomain, err := (*repository.ColumnDS()).Create(ctx, input.title, input.position, input.boardId)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Column: *columnDomain}, nil
}

type CreateInput struct {
	title       string
	position    float64
	boardId     int
	firebaseUID string
}

func NewCreateInput(title string, position float64, boardId int, firebaseUID string) *CreateInput {
	return &CreateInput{title: title, position: position, boardId: boardId, firebaseUID: firebaseUID}
}

type CreateOutput struct {
	Column column.Column
}
