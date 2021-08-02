package datasource

import (
	"context"
	"drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastracture/mysql"
	"fmt"
)

type Workspace struct{}

func (w Workspace) ListWorkspaces(ctx context.Context) (*[]*workspace.Workspace, error) {
	ws, err := mysql.Client().Workspace.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	workspaces := []*workspace.Workspace{}
	for _, w := range ws {
		workspaces = append(workspaces, workspace.New(w.ID, w.Title))
	}

	return &workspaces, nil
}
