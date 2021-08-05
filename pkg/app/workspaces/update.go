package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/workspace"
)

func Update(ctx context.Context, workspaceRepo repository.Workspace, input *UpdateInput) (*UpdateOutput, error) {
	wdomain, err := workspaceRepo.Update(ctx, input.ID, input.Title)
	if err != nil {
		return nil, err
	}

	return &UpdateOutput{Workspace: *wdomain}, nil
}

type UpdateInput struct {
	ID    int
	Title string
}

type UpdateOutput struct {
	Workspace workspace.Workspace
}
