package repository

import (
	"context"
	"drello-api/pkg/domain/workspace"
)

type Workspace interface {
	ListWorkspaces(ctx context.Context) (*[]*workspace.Workspace, error)
}
