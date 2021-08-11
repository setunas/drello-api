package datasource

import (
	"context"
	"drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastracture/mysql"
	"fmt"
)

type Workspace struct{}

func (w Workspace) List(ctx context.Context) (*[]*workspace.Workspace, error) {
	ws, err := mysql.Client().Workspace.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying workspace: %w", err)
	}

	workspaces := []*workspace.Workspace{}
	for _, w := range ws {
		workspaces = append(workspaces, workspace.New(w.ID, w.Title))
	}

	return &workspaces, nil
}

func (w Workspace) GetOne(ctx context.Context, id int) (*workspace.Workspace, error) {
	wNode, err := mysql.Client().Workspace.Query().Where(eWorkspace.ID(id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying workspace: %w", err)
	}

	return workspace.New(wNode.ID, wNode.Title), nil
}

func (w Workspace) Create(ctx context.Context, title string) (*workspace.Workspace, error) {
	wNode, err := mysql.Client().Workspace.Create().SetTitle(title).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating workspace: %w", err)
	}

	return workspace.New(wNode.ID, wNode.Title), nil
}

func (w Workspace) Update(ctx context.Context, id int, title string) (*workspace.Workspace, error) {
	wNode, err := mysql.Client().Workspace.UpdateOneID(id).SetTitle(title).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating workspace: %w", err)
	}

	return workspace.New(wNode.ID, wNode.Title), nil
}

func (w Workspace) Delete(ctx context.Context, id int) error {
	err := mysql.Client().Workspace.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed deleting workspace: %w", err)
	}

	return nil
}
