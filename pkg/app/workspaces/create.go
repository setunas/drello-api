package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/workspace"
)

func Create(ctx context.Context, workspaceRepo repository.Workspace, input *CreateInput) (*CreateOutput, error) {
	wdomain, err := workspaceRepo.Create(ctx, input.Title)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Workspace: *wdomain}, nil
}

type CreateInput struct {
	Title string
}

type CreateOutput struct {
	Workspace workspace.Workspace
}
