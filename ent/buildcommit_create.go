// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/buildcommit"
	"github.com/gen0cide/laforge/ent/plandiff"
	"github.com/gen0cide/laforge/ent/servertask"
	"github.com/google/uuid"
)

// BuildCommitCreate is the builder for creating a BuildCommit entity.
type BuildCommitCreate struct {
	config
	mutation *BuildCommitMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (bcc *BuildCommitCreate) SetType(b buildcommit.Type) *BuildCommitCreate {
	bcc.mutation.SetType(b)
	return bcc
}

// SetRevision sets the "revision" field.
func (bcc *BuildCommitCreate) SetRevision(i int) *BuildCommitCreate {
	bcc.mutation.SetRevision(i)
	return bcc
}

// SetState sets the "state" field.
func (bcc *BuildCommitCreate) SetState(b buildcommit.State) *BuildCommitCreate {
	bcc.mutation.SetState(b)
	return bcc
}

// SetCreatedAt sets the "created_at" field.
func (bcc *BuildCommitCreate) SetCreatedAt(t time.Time) *BuildCommitCreate {
	bcc.mutation.SetCreatedAt(t)
	return bcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bcc *BuildCommitCreate) SetNillableCreatedAt(t *time.Time) *BuildCommitCreate {
	if t != nil {
		bcc.SetCreatedAt(*t)
	}
	return bcc
}

// SetID sets the "id" field.
func (bcc *BuildCommitCreate) SetID(u uuid.UUID) *BuildCommitCreate {
	bcc.mutation.SetID(u)
	return bcc
}

// SetBuildCommitToBuildID sets the "BuildCommitToBuild" edge to the Build entity by ID.
func (bcc *BuildCommitCreate) SetBuildCommitToBuildID(id uuid.UUID) *BuildCommitCreate {
	bcc.mutation.SetBuildCommitToBuildID(id)
	return bcc
}

// SetBuildCommitToBuild sets the "BuildCommitToBuild" edge to the Build entity.
func (bcc *BuildCommitCreate) SetBuildCommitToBuild(b *Build) *BuildCommitCreate {
	return bcc.SetBuildCommitToBuildID(b.ID)
}

// AddBuildCommitToServerTaskIDs adds the "BuildCommitToServerTask" edge to the ServerTask entity by IDs.
func (bcc *BuildCommitCreate) AddBuildCommitToServerTaskIDs(ids ...uuid.UUID) *BuildCommitCreate {
	bcc.mutation.AddBuildCommitToServerTaskIDs(ids...)
	return bcc
}

// AddBuildCommitToServerTask adds the "BuildCommitToServerTask" edges to the ServerTask entity.
func (bcc *BuildCommitCreate) AddBuildCommitToServerTask(s ...*ServerTask) *BuildCommitCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bcc.AddBuildCommitToServerTaskIDs(ids...)
}

// AddBuildCommitToPlanDiffIDs adds the "BuildCommitToPlanDiffs" edge to the PlanDiff entity by IDs.
func (bcc *BuildCommitCreate) AddBuildCommitToPlanDiffIDs(ids ...uuid.UUID) *BuildCommitCreate {
	bcc.mutation.AddBuildCommitToPlanDiffIDs(ids...)
	return bcc
}

// AddBuildCommitToPlanDiffs adds the "BuildCommitToPlanDiffs" edges to the PlanDiff entity.
func (bcc *BuildCommitCreate) AddBuildCommitToPlanDiffs(p ...*PlanDiff) *BuildCommitCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bcc.AddBuildCommitToPlanDiffIDs(ids...)
}

// Mutation returns the BuildCommitMutation object of the builder.
func (bcc *BuildCommitCreate) Mutation() *BuildCommitMutation {
	return bcc.mutation
}

// Save creates the BuildCommit in the database.
func (bcc *BuildCommitCreate) Save(ctx context.Context) (*BuildCommit, error) {
	var (
		err  error
		node *BuildCommit
	)
	bcc.defaults()
	if len(bcc.hooks) == 0 {
		if err = bcc.check(); err != nil {
			return nil, err
		}
		node, err = bcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcc.check(); err != nil {
				return nil, err
			}
			bcc.mutation = mutation
			if node, err = bcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bcc.hooks) - 1; i >= 0; i-- {
			if bcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bcc *BuildCommitCreate) SaveX(ctx context.Context) *BuildCommit {
	v, err := bcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcc *BuildCommitCreate) Exec(ctx context.Context) error {
	_, err := bcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcc *BuildCommitCreate) ExecX(ctx context.Context) {
	if err := bcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcc *BuildCommitCreate) defaults() {
	if _, ok := bcc.mutation.CreatedAt(); !ok {
		v := buildcommit.DefaultCreatedAt()
		bcc.mutation.SetCreatedAt(v)
	}
	if _, ok := bcc.mutation.ID(); !ok {
		v := buildcommit.DefaultID()
		bcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcc *BuildCommitCreate) check() error {
	if _, ok := bcc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := bcc.mutation.GetType(); ok {
		if err := buildcommit.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := bcc.mutation.Revision(); !ok {
		return &ValidationError{Name: "revision", err: errors.New(`ent: missing required field "revision"`)}
	}
	if _, ok := bcc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := bcc.mutation.State(); ok {
		if err := buildcommit.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := bcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := bcc.mutation.BuildCommitToBuildID(); !ok {
		return &ValidationError{Name: "BuildCommitToBuild", err: errors.New("ent: missing required edge \"BuildCommitToBuild\"")}
	}
	return nil
}

func (bcc *BuildCommitCreate) sqlSave(ctx context.Context) (*BuildCommit, error) {
	_node, _spec := bcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (bcc *BuildCommitCreate) createSpec() (*BuildCommit, *sqlgraph.CreateSpec) {
	var (
		_node = &BuildCommit{config: bcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: buildcommit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: buildcommit.FieldID,
			},
		}
	)
	if id, ok := bcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bcc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: buildcommit.FieldType,
		})
		_node.Type = value
	}
	if value, ok := bcc.mutation.Revision(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: buildcommit.FieldRevision,
		})
		_node.Revision = value
	}
	if value, ok := bcc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: buildcommit.FieldState,
		})
		_node.State = value
	}
	if value, ok := bcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: buildcommit.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := bcc.mutation.BuildCommitToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   buildcommit.BuildCommitToBuildTable,
			Columns: []string{buildcommit.BuildCommitToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.build_commit_build_commit_to_build = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.BuildCommitToServerTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToServerTaskTable,
			Columns: []string{buildcommit.BuildCommitToServerTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servertask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.BuildCommitToPlanDiffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BuildCommitCreateBulk is the builder for creating many BuildCommit entities in bulk.
type BuildCommitCreateBulk struct {
	config
	builders []*BuildCommitCreate
}

// Save creates the BuildCommit entities in the database.
func (bccb *BuildCommitCreateBulk) Save(ctx context.Context) ([]*BuildCommit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bccb.builders))
	nodes := make([]*BuildCommit, len(bccb.builders))
	mutators := make([]Mutator, len(bccb.builders))
	for i := range bccb.builders {
		func(i int, root context.Context) {
			builder := bccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BuildCommitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bccb *BuildCommitCreateBulk) SaveX(ctx context.Context) []*BuildCommit {
	v, err := bccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bccb *BuildCommitCreateBulk) Exec(ctx context.Context) error {
	_, err := bccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bccb *BuildCommitCreateBulk) ExecX(ctx context.Context) {
	if err := bccb.Exec(ctx); err != nil {
		panic(err)
	}
}
