package datasource

import (
	"context"
	"database/sql"
	"drello-api/pkg/domain/workspace"
	domainWorkspace "drello-api/pkg/domain/workspace"
	"drello-api/pkg/infrastracture/mysql"
	"fmt"
)

type Workspace struct{}

func (w Workspace) GetAll(ctx context.Context) (*[]*domainWorkspace.Workspace, error) {
	db := mysql.DBPool()
	rows, err := db.Query("SELECT * FROM workspaces")
	if err != nil {
		return nil, fmt.Errorf("failed querying workspaces: %w", err)
	}
	defer rows.Close()

	workspaces := []*domainWorkspace.Workspace{}

	for rows.Next() {
		var id int
		var title string

		err := rows.Scan(&id, &title)
		if err != nil {
			return nil, fmt.Errorf("failed querying workspaces: %w", err)
		}

		workspaces = append(workspaces, domainWorkspace.New(id, title))
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed querying workspaces: %w", err)
	}

	return &workspaces, nil
}

func (w Workspace) GetOne(ctx context.Context, id int) (*domainWorkspace.Workspace, error) {
	var title string

	db := mysql.DBPool()
	row := db.QueryRow("SELECT id, title FROM workspaces WHERE id = ?", id)

	switch err := row.Scan(&title); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with id %d", id)
	case nil:
		return domainWorkspace.New(id, title), nil
	default:
		return nil, err
	}
}

func (w Workspace) Create(ctx context.Context, title string) (*domainWorkspace.Workspace, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO workspaces (title) VALUES (?)", title)
	if err != nil {
		return nil, fmt.Errorf("failed creating workspace: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating workspace: %w", err)
	}

	return workspace.New(int(lastInsertID), title), nil
}

func (w Workspace) Update(ctx context.Context, id int, title string) (*domainWorkspace.Workspace, error) {
	db := mysql.DBPool()
	_, err := db.Exec("UPDATE workspaces SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating workspace: %w", err)
	}

	return workspace.New(id, title), nil
}

func (w Workspace) Delete(ctx context.Context, id int) error {
	db := mysql.DBPool()
	_, err := db.Exec("DELETE FROM workspaces WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed deleting workspace: %w", err)
	}

	return nil
}
