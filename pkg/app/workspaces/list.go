package workspaces

import (
	"drello-api/pkg/app/repository"
)

func List(workspaceRepo repository.Workspace) *ListOutput {
	workspaces := workspaceRepo.ListWorkspaces()

	titles := []string{}

	for _, w := range *workspaces {
		titles = append(titles, w.Title())
	}

	return &ListOutput{titles: titles}
}

type ListOutput struct {
	titles []string
}

func (lo ListOutput) Titles() []string {
	return lo.titles
}
