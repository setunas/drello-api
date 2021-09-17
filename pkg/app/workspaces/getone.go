package workspaces

import (
	"context"
	wdomain "drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastructure/repository"
)

func GetOne(ctx context.Context, workspaceRepo repository.Workspace, input *GetOneInput) (*GetOneOutput, error) {
	wNode, err := workspaceRepo.GetOne(ctx, input.id)
	if err != nil {
		return nil, err
	}

	return &GetOneOutput{Workspace: wdomain.New(wNode.ID(), wNode.Title())}, nil
}

type GetOneInput struct {
	id int
}

func NewGetOneInput(id int) *GetOneInput {
	return &GetOneInput{id: id}
}

type GetOneOutput struct {
	Workspace *wdomain.Workspace
}
