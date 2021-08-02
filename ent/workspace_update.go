// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"drello-api/ent/predicate"
	"drello-api/ent/workspace"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceUpdate is the builder for updating Workspace entities.
type WorkspaceUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// Where adds a new predicate for the WorkspaceUpdate builder.
func (wu *WorkspaceUpdate) Where(ps ...predicate.Workspace) *WorkspaceUpdate {
	wu.mutation.predicates = append(wu.mutation.predicates, ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WorkspaceUpdate) SetTitle(s string) *WorkspaceUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wu *WorkspaceUpdate) Mutation() *WorkspaceMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkspaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkspaceUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkspaceUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkspaceUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wu *WorkspaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workspace.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldTitle,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceUpdateOne is the builder for updating a single Workspace entity.
type WorkspaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkspaceMutation
}

// SetTitle sets the "title" field.
func (wuo *WorkspaceUpdateOne) SetTitle(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wuo *WorkspaceUpdateOne) Mutation() *WorkspaceMutation {
	return wuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkspaceUpdateOne) Select(field string, fields ...string) *WorkspaceUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workspace entity.
func (wuo *WorkspaceUpdateOne) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) SaveX(ctx context.Context) *Workspace {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkspaceUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuo *WorkspaceUpdateOne) sqlSave(ctx context.Context) (_node *Workspace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workspace.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Workspace.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspace.FieldID)
		for _, f := range fields {
			if !workspace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldTitle,
		})
	}
	_node = &Workspace{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
