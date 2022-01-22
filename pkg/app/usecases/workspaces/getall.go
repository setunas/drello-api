package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	workspaceDomain "drello-api/pkg/domain/workspace"
)

func GetAll(ctx context.Context, workspaceRepo repository.Workspace) (*GetAllOutput, error) {
	workspaces, err := workspaceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var workspaceDomains []*workspaceDomain.Workspace

	for _, w := range *workspaces {
		workspaceDomains = append(workspaceDomains, workspaceDomain.New(w.ID(), w.Title()))
	}

	return &GetAllOutput{Workspaces: workspaceDomains}, nil
}

type GetAllOutput struct {
	Workspaces []*workspaceDomain.Workspace
}
