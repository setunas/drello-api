package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/workspace"
)

func Create(ctx context.Context, workspaceRepo repository.Workspace, input *CreateInput) (*CreateOutput, error) {
	wdomain, err := workspaceRepo.Create(ctx, input.title)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{Workspace: *wdomain}, nil
}

type CreateInput struct {
	title string
}

func NewCreateInput(title string) *CreateInput {
	return &CreateInput{title: title}
}

type CreateOutput struct {
	Workspace workspace.Workspace
}
