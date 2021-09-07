package workspaces

import (
	"context"
	wdomain "drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastructure/repository"
)

func GetAll(ctx context.Context, workspaceRepo repository.Workspace) (*GetAllOutput, error) {
	wNodes, err := workspaceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var wdomains []*wdomain.Workspace

	for _, w := range *wNodes {
		wdomains = append(wdomains, wdomain.New(w.ID(), w.Title()))
	}

	return &GetAllOutput{Workspaces: wdomains}, nil
}

type GetAllOutput struct {
	Workspaces []*wdomain.Workspace
}
