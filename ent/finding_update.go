// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/finding"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// FindingUpdate is the builder for updating Finding entities.
type FindingUpdate struct {
	config
	hooks    []Hook
	mutation *FindingMutation
}

// Where adds a new predicate for the builder.
func (fu *FindingUpdate) Where(ps ...predicate.Finding) *FindingUpdate {
	fu.mutation.predicates = append(fu.mutation.predicates, ps...)
	return fu
}

// SetName sets the name field.
func (fu *FindingUpdate) SetName(s string) *FindingUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetDescription sets the description field.
func (fu *FindingUpdate) SetDescription(s string) *FindingUpdate {
	fu.mutation.SetDescription(s)
	return fu
}

// SetSeverity sets the severity field.
func (fu *FindingUpdate) SetSeverity(f finding.Severity) *FindingUpdate {
	fu.mutation.SetSeverity(f)
	return fu
}

// SetDifficulty sets the difficulty field.
func (fu *FindingUpdate) SetDifficulty(f finding.Difficulty) *FindingUpdate {
	fu.mutation.SetDifficulty(f)
	return fu
}

// AddUserIDs adds the user edge to User by ids.
func (fu *FindingUpdate) AddUserIDs(ids ...int) *FindingUpdate {
	fu.mutation.AddUserIDs(ids...)
	return fu
}

// AddUser adds the user edges to User.
func (fu *FindingUpdate) AddUser(u ...*User) *FindingUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fu.AddUserIDs(ids...)
}

// AddTagIDs adds the tag edge to Tag by ids.
func (fu *FindingUpdate) AddTagIDs(ids ...int) *FindingUpdate {
	fu.mutation.AddTagIDs(ids...)
	return fu
}

// AddTag adds the tag edges to Tag.
func (fu *FindingUpdate) AddTag(t ...*Tag) *FindingUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fu.AddTagIDs(ids...)
}

// AddHostIDs adds the host edge to Host by ids.
func (fu *FindingUpdate) AddHostIDs(ids ...int) *FindingUpdate {
	fu.mutation.AddHostIDs(ids...)
	return fu
}

// AddHost adds the host edges to Host.
func (fu *FindingUpdate) AddHost(h ...*Host) *FindingUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return fu.AddHostIDs(ids...)
}

// Mutation returns the FindingMutation object of the builder.
func (fu *FindingUpdate) Mutation() *FindingMutation {
	return fu.mutation
}

// ClearUser clears all "user" edges to type User.
func (fu *FindingUpdate) ClearUser() *FindingUpdate {
	fu.mutation.ClearUser()
	return fu
}

// RemoveUserIDs removes the user edge to User by ids.
func (fu *FindingUpdate) RemoveUserIDs(ids ...int) *FindingUpdate {
	fu.mutation.RemoveUserIDs(ids...)
	return fu
}

// RemoveUser removes user edges to User.
func (fu *FindingUpdate) RemoveUser(u ...*User) *FindingUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fu.RemoveUserIDs(ids...)
}

// ClearTag clears all "tag" edges to type Tag.
func (fu *FindingUpdate) ClearTag() *FindingUpdate {
	fu.mutation.ClearTag()
	return fu
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (fu *FindingUpdate) RemoveTagIDs(ids ...int) *FindingUpdate {
	fu.mutation.RemoveTagIDs(ids...)
	return fu
}

// RemoveTag removes tag edges to Tag.
func (fu *FindingUpdate) RemoveTag(t ...*Tag) *FindingUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fu.RemoveTagIDs(ids...)
}

// ClearHost clears all "host" edges to type Host.
func (fu *FindingUpdate) ClearHost() *FindingUpdate {
	fu.mutation.ClearHost()
	return fu
}

// RemoveHostIDs removes the host edge to Host by ids.
func (fu *FindingUpdate) RemoveHostIDs(ids ...int) *FindingUpdate {
	fu.mutation.RemoveHostIDs(ids...)
	return fu
}

// RemoveHost removes host edges to Host.
func (fu *FindingUpdate) RemoveHost(h ...*Host) *FindingUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return fu.RemoveHostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FindingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FindingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FindingUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FindingUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FindingUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FindingUpdate) check() error {
	if v, ok := fu.mutation.Severity(); ok {
		if err := finding.SeverityValidator(v); err != nil {
			return &ValidationError{Name: "severity", err: fmt.Errorf("ent: validator failed for field \"severity\": %w", err)}
		}
	}
	if v, ok := fu.mutation.Difficulty(); ok {
		if err := finding.DifficultyValidator(v); err != nil {
			return &ValidationError{Name: "difficulty", err: fmt.Errorf("ent: validator failed for field \"difficulty\": %w", err)}
		}
	}
	return nil
}

func (fu *FindingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   finding.Table,
			Columns: finding.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: finding.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: finding.FieldName,
		})
	}
	if value, ok := fu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: finding.FieldDescription,
		})
	}
	if value, ok := fu.mutation.Severity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: finding.FieldSeverity,
		})
	}
	if value, ok := fu.mutation.Difficulty(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: finding.FieldDifficulty,
		})
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.UserTable,
			Columns: []string{finding.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedUserIDs(); len(nodes) > 0 && !fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.UserTable,
			Columns: []string{finding.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.UserTable,
			Columns: []string{finding.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.TagTable,
			Columns: []string{finding.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedTagIDs(); len(nodes) > 0 && !fu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.TagTable,
			Columns: []string{finding.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.TagTable,
			Columns: []string{finding.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.HostTable,
			Columns: []string{finding.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedHostIDs(); len(nodes) > 0 && !fu.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.HostTable,
			Columns: []string{finding.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.HostTable,
			Columns: []string{finding.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{finding.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FindingUpdateOne is the builder for updating a single Finding entity.
type FindingUpdateOne struct {
	config
	hooks    []Hook
	mutation *FindingMutation
}

// SetName sets the name field.
func (fuo *FindingUpdateOne) SetName(s string) *FindingUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetDescription sets the description field.
func (fuo *FindingUpdateOne) SetDescription(s string) *FindingUpdateOne {
	fuo.mutation.SetDescription(s)
	return fuo
}

// SetSeverity sets the severity field.
func (fuo *FindingUpdateOne) SetSeverity(f finding.Severity) *FindingUpdateOne {
	fuo.mutation.SetSeverity(f)
	return fuo
}

// SetDifficulty sets the difficulty field.
func (fuo *FindingUpdateOne) SetDifficulty(f finding.Difficulty) *FindingUpdateOne {
	fuo.mutation.SetDifficulty(f)
	return fuo
}

// AddUserIDs adds the user edge to User by ids.
func (fuo *FindingUpdateOne) AddUserIDs(ids ...int) *FindingUpdateOne {
	fuo.mutation.AddUserIDs(ids...)
	return fuo
}

// AddUser adds the user edges to User.
func (fuo *FindingUpdateOne) AddUser(u ...*User) *FindingUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fuo.AddUserIDs(ids...)
}

// AddTagIDs adds the tag edge to Tag by ids.
func (fuo *FindingUpdateOne) AddTagIDs(ids ...int) *FindingUpdateOne {
	fuo.mutation.AddTagIDs(ids...)
	return fuo
}

// AddTag adds the tag edges to Tag.
func (fuo *FindingUpdateOne) AddTag(t ...*Tag) *FindingUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fuo.AddTagIDs(ids...)
}

// AddHostIDs adds the host edge to Host by ids.
func (fuo *FindingUpdateOne) AddHostIDs(ids ...int) *FindingUpdateOne {
	fuo.mutation.AddHostIDs(ids...)
	return fuo
}

// AddHost adds the host edges to Host.
func (fuo *FindingUpdateOne) AddHost(h ...*Host) *FindingUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return fuo.AddHostIDs(ids...)
}

// Mutation returns the FindingMutation object of the builder.
func (fuo *FindingUpdateOne) Mutation() *FindingMutation {
	return fuo.mutation
}

// ClearUser clears all "user" edges to type User.
func (fuo *FindingUpdateOne) ClearUser() *FindingUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// RemoveUserIDs removes the user edge to User by ids.
func (fuo *FindingUpdateOne) RemoveUserIDs(ids ...int) *FindingUpdateOne {
	fuo.mutation.RemoveUserIDs(ids...)
	return fuo
}

// RemoveUser removes user edges to User.
func (fuo *FindingUpdateOne) RemoveUser(u ...*User) *FindingUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fuo.RemoveUserIDs(ids...)
}

// ClearTag clears all "tag" edges to type Tag.
func (fuo *FindingUpdateOne) ClearTag() *FindingUpdateOne {
	fuo.mutation.ClearTag()
	return fuo
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (fuo *FindingUpdateOne) RemoveTagIDs(ids ...int) *FindingUpdateOne {
	fuo.mutation.RemoveTagIDs(ids...)
	return fuo
}

// RemoveTag removes tag edges to Tag.
func (fuo *FindingUpdateOne) RemoveTag(t ...*Tag) *FindingUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fuo.RemoveTagIDs(ids...)
}

// ClearHost clears all "host" edges to type Host.
func (fuo *FindingUpdateOne) ClearHost() *FindingUpdateOne {
	fuo.mutation.ClearHost()
	return fuo
}

// RemoveHostIDs removes the host edge to Host by ids.
func (fuo *FindingUpdateOne) RemoveHostIDs(ids ...int) *FindingUpdateOne {
	fuo.mutation.RemoveHostIDs(ids...)
	return fuo
}

// RemoveHost removes host edges to Host.
func (fuo *FindingUpdateOne) RemoveHost(h ...*Host) *FindingUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return fuo.RemoveHostIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (fuo *FindingUpdateOne) Save(ctx context.Context) (*Finding, error) {
	var (
		err  error
		node *Finding
	)
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FindingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FindingUpdateOne) SaveX(ctx context.Context) *Finding {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FindingUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FindingUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FindingUpdateOne) check() error {
	if v, ok := fuo.mutation.Severity(); ok {
		if err := finding.SeverityValidator(v); err != nil {
			return &ValidationError{Name: "severity", err: fmt.Errorf("ent: validator failed for field \"severity\": %w", err)}
		}
	}
	if v, ok := fuo.mutation.Difficulty(); ok {
		if err := finding.DifficultyValidator(v); err != nil {
			return &ValidationError{Name: "difficulty", err: fmt.Errorf("ent: validator failed for field \"difficulty\": %w", err)}
		}
	}
	return nil
}

func (fuo *FindingUpdateOne) sqlSave(ctx context.Context) (_node *Finding, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   finding.Table,
			Columns: finding.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: finding.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Finding.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: finding.FieldName,
		})
	}
	if value, ok := fuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: finding.FieldDescription,
		})
	}
	if value, ok := fuo.mutation.Severity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: finding.FieldSeverity,
		})
	}
	if value, ok := fuo.mutation.Difficulty(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: finding.FieldDifficulty,
		})
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.UserTable,
			Columns: []string{finding.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedUserIDs(); len(nodes) > 0 && !fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.UserTable,
			Columns: []string{finding.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.UserTable,
			Columns: []string{finding.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.TagTable,
			Columns: []string{finding.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedTagIDs(); len(nodes) > 0 && !fuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.TagTable,
			Columns: []string{finding.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.TagTable,
			Columns: []string{finding.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.HostTable,
			Columns: []string{finding.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedHostIDs(); len(nodes) > 0 && !fuo.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.HostTable,
			Columns: []string{finding.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   finding.HostTable,
			Columns: []string{finding.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Finding{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{finding.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
