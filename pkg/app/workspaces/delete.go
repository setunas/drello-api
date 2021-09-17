package workspaces

import (
	"context"
	"drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastructure/repository"
)

func Delete(ctx context.Context, workspaceRepo repository.Workspace, input *DeleteInput) error {
	err := workspaceRepo.Delete(ctx, input.id)
	if err != nil {
		return err
	}

	return nil
}

type DeleteInput struct {
	id int
}

func NewDeleteInput(id int) *DeleteInput {
	return &DeleteInput{id: id}
}

type DeleteOutput struct {
	Workspace workspace.Workspace
}
