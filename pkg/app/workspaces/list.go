package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
)

func List(ctx context.Context, workspaceRepo repository.Workspace) (*ListOutput, error) {
	workspaces, err := workspaceRepo.ListWorkspaces(ctx)
	if err != nil {
		return nil, err
	}

	titles := []string{}

	for _, w := range *workspaces {
		titles = append(titles, w.Title())
	}

	return &ListOutput{titles: titles}, nil
}

type ListOutput struct {
	titles []string
}

func (lo ListOutput) Titles() []string {
	return lo.titles
}
