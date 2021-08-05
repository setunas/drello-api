package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
	wdomain "drello-api/pkg/domain/workspace"
)

func List(ctx context.Context, workspaceRepo repository.Workspace) (*ListOutput, error) {
	wNodes, err := workspaceRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var wdomains []*wdomain.Workspace

	for _, w := range *wNodes {
		wdomains = append(wdomains, wdomain.New(w.ID(), w.Title()))
	}

	return &ListOutput{Workspaces: wdomains}, nil
}

type ListOutput struct {
	Workspaces []*wdomain.Workspace
}
