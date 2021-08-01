package repository

import "drello-api/pkg/domain/workspace"

type Workspace interface {
	ListWorkspaces() *[]*workspace.Workspace
}
