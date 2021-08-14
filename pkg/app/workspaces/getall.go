package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	wdomain "drello-api/pkg/domain/workspace"
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
