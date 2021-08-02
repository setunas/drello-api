package workspaces

import (
	"context"
	"drello-api/pkg/app/repository"
)

func List(ctx context.Context, workspaceRepo repository.Workspace) (*ListOutput, error) {
	workspaces, err := workspaceRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	titles := []string{}

	for _, w := range *workspaces {
		titles = append(titles, w.Title())
	}

	return &ListOutput{Titles: titles}, nil
}

type ListOutput struct {
	Titles []string
}
