package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	workspaceDomain "drello-api/pkg/domain/workspace"
)

func GetOne(ctx context.Context, workspaceRepo repository.Workspace, input *GetOneInput) (*GetOneOutput, error) {
	workspaces, err := workspaceRepo.GetOne(ctx, input.id)
	if err != nil {
		return nil, err
	}

	return &GetOneOutput{Workspace: workspaceDomain.New(workspaces.ID(), workspaces.Title())}, nil
}

type GetOneInput struct {
	id int
}

func NewGetOneInput(id int) *GetOneInput {
	return &GetOneInput{id: id}
}

type GetOneOutput struct {
	Workspace *workspaceDomain.Workspace
}
