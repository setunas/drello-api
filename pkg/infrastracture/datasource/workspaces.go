package datasource

import (
	"drello-api/pkg/domain/workspace"
)

type Workspace struct{}

func (w Workspace) ListWorkspaces() *[]*workspace.Workspace {

	workspaces := []*workspace.Workspace{}
	workspaces = append(workspaces, workspace.New(1, "one"))
	workspaces = append(workspaces, workspace.New(2, "two"))
	workspaces = append(workspaces, workspace.New(3, "three"))

	return &workspaces
}
