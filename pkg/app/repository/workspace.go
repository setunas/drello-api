package repository

import (
	"context"
	"drello-api/pkg/domain/workspace"
)

type Workspace interface {
	List(ctx context.Context) (*[]*workspace.Workspace, error)
	Create(ctx context.Context, title string) (*workspace.Workspace, error)
	Update(ctx context.Context, id int, title string) (*workspace.Workspace, error)
	Delete(ctx context.Context, id int) error
}
