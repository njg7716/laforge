// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/repocommit"
	"github.com/gen0cide/laforge/ent/repository"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/uuid"
)

// RepoCommitUpdate is the builder for updating RepoCommit entities.
type RepoCommitUpdate struct {
	config
	hooks    []Hook
	mutation *RepoCommitMutation
}

// Where appends a list predicates to the RepoCommitUpdate builder.
func (rcu *RepoCommitUpdate) Where(ps ...predicate.RepoCommit) *RepoCommitUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetRevision sets the "revision" field.
func (rcu *RepoCommitUpdate) SetRevision(i int) *RepoCommitUpdate {
	rcu.mutation.ResetRevision()
	rcu.mutation.SetRevision(i)
	return rcu
}

// AddRevision adds i to the "revision" field.
func (rcu *RepoCommitUpdate) AddRevision(i int) *RepoCommitUpdate {
	rcu.mutation.AddRevision(i)
	return rcu
}

// SetHash sets the "hash" field.
func (rcu *RepoCommitUpdate) SetHash(s string) *RepoCommitUpdate {
	rcu.mutation.SetHash(s)
	return rcu
}

// SetAuthor sets the "author" field.
func (rcu *RepoCommitUpdate) SetAuthor(o object.Signature) *RepoCommitUpdate {
	rcu.mutation.SetAuthor(o)
	return rcu
}

// SetCommitter sets the "committer" field.
func (rcu *RepoCommitUpdate) SetCommitter(o object.Signature) *RepoCommitUpdate {
	rcu.mutation.SetCommitter(o)
	return rcu
}

// SetPgpSignature sets the "pgp_signature" field.
func (rcu *RepoCommitUpdate) SetPgpSignature(s string) *RepoCommitUpdate {
	rcu.mutation.SetPgpSignature(s)
	return rcu
}

// SetMessage sets the "message" field.
func (rcu *RepoCommitUpdate) SetMessage(s string) *RepoCommitUpdate {
	rcu.mutation.SetMessage(s)
	return rcu
}

// SetTreeHash sets the "tree_hash" field.
func (rcu *RepoCommitUpdate) SetTreeHash(s string) *RepoCommitUpdate {
	rcu.mutation.SetTreeHash(s)
	return rcu
}

// SetParentHashes sets the "parent_hashes" field.
func (rcu *RepoCommitUpdate) SetParentHashes(s []string) *RepoCommitUpdate {
	rcu.mutation.SetParentHashes(s)
	return rcu
}

// SetRepoCommitToRepositoryID sets the "RepoCommitToRepository" edge to the Repository entity by ID.
func (rcu *RepoCommitUpdate) SetRepoCommitToRepositoryID(id uuid.UUID) *RepoCommitUpdate {
	rcu.mutation.SetRepoCommitToRepositoryID(id)
	return rcu
}

// SetNillableRepoCommitToRepositoryID sets the "RepoCommitToRepository" edge to the Repository entity by ID if the given value is not nil.
func (rcu *RepoCommitUpdate) SetNillableRepoCommitToRepositoryID(id *uuid.UUID) *RepoCommitUpdate {
	if id != nil {
		rcu = rcu.SetRepoCommitToRepositoryID(*id)
	}
	return rcu
}

// SetRepoCommitToRepository sets the "RepoCommitToRepository" edge to the Repository entity.
func (rcu *RepoCommitUpdate) SetRepoCommitToRepository(r *Repository) *RepoCommitUpdate {
	return rcu.SetRepoCommitToRepositoryID(r.ID)
}

// Mutation returns the RepoCommitMutation object of the builder.
func (rcu *RepoCommitUpdate) Mutation() *RepoCommitMutation {
	return rcu.mutation
}

// ClearRepoCommitToRepository clears the "RepoCommitToRepository" edge to the Repository entity.
func (rcu *RepoCommitUpdate) ClearRepoCommitToRepository() *RepoCommitUpdate {
	rcu.mutation.ClearRepoCommitToRepository()
	return rcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *RepoCommitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rcu.hooks) == 0 {
		affected, err = rcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rcu.mutation = mutation
			affected, err = rcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rcu.hooks) - 1; i >= 0; i-- {
			if rcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *RepoCommitUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *RepoCommitUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *RepoCommitUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rcu *RepoCommitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repocommit.Table,
			Columns: repocommit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: repocommit.FieldID,
			},
		},
	}
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repocommit.FieldRevision,
		})
	}
	if value, ok := rcu.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repocommit.FieldRevision,
		})
	}
	if value, ok := rcu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldHash,
		})
	}
	if value, ok := rcu.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: repocommit.FieldAuthor,
		})
	}
	if value, ok := rcu.mutation.Committer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: repocommit.FieldCommitter,
		})
	}
	if value, ok := rcu.mutation.PgpSignature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldPgpSignature,
		})
	}
	if value, ok := rcu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldMessage,
		})
	}
	if value, ok := rcu.mutation.TreeHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldTreeHash,
		})
	}
	if value, ok := rcu.mutation.ParentHashes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: repocommit.FieldParentHashes,
		})
	}
	if rcu.mutation.RepoCommitToRepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repocommit.RepoCommitToRepositoryTable,
			Columns: []string{repocommit.RepoCommitToRepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcu.mutation.RepoCommitToRepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repocommit.RepoCommitToRepositoryTable,
			Columns: []string{repocommit.RepoCommitToRepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repocommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RepoCommitUpdateOne is the builder for updating a single RepoCommit entity.
type RepoCommitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepoCommitMutation
}

// SetRevision sets the "revision" field.
func (rcuo *RepoCommitUpdateOne) SetRevision(i int) *RepoCommitUpdateOne {
	rcuo.mutation.ResetRevision()
	rcuo.mutation.SetRevision(i)
	return rcuo
}

// AddRevision adds i to the "revision" field.
func (rcuo *RepoCommitUpdateOne) AddRevision(i int) *RepoCommitUpdateOne {
	rcuo.mutation.AddRevision(i)
	return rcuo
}

// SetHash sets the "hash" field.
func (rcuo *RepoCommitUpdateOne) SetHash(s string) *RepoCommitUpdateOne {
	rcuo.mutation.SetHash(s)
	return rcuo
}

// SetAuthor sets the "author" field.
func (rcuo *RepoCommitUpdateOne) SetAuthor(o object.Signature) *RepoCommitUpdateOne {
	rcuo.mutation.SetAuthor(o)
	return rcuo
}

// SetCommitter sets the "committer" field.
func (rcuo *RepoCommitUpdateOne) SetCommitter(o object.Signature) *RepoCommitUpdateOne {
	rcuo.mutation.SetCommitter(o)
	return rcuo
}

// SetPgpSignature sets the "pgp_signature" field.
func (rcuo *RepoCommitUpdateOne) SetPgpSignature(s string) *RepoCommitUpdateOne {
	rcuo.mutation.SetPgpSignature(s)
	return rcuo
}

// SetMessage sets the "message" field.
func (rcuo *RepoCommitUpdateOne) SetMessage(s string) *RepoCommitUpdateOne {
	rcuo.mutation.SetMessage(s)
	return rcuo
}

// SetTreeHash sets the "tree_hash" field.
func (rcuo *RepoCommitUpdateOne) SetTreeHash(s string) *RepoCommitUpdateOne {
	rcuo.mutation.SetTreeHash(s)
	return rcuo
}

// SetParentHashes sets the "parent_hashes" field.
func (rcuo *RepoCommitUpdateOne) SetParentHashes(s []string) *RepoCommitUpdateOne {
	rcuo.mutation.SetParentHashes(s)
	return rcuo
}

// SetRepoCommitToRepositoryID sets the "RepoCommitToRepository" edge to the Repository entity by ID.
func (rcuo *RepoCommitUpdateOne) SetRepoCommitToRepositoryID(id uuid.UUID) *RepoCommitUpdateOne {
	rcuo.mutation.SetRepoCommitToRepositoryID(id)
	return rcuo
}

// SetNillableRepoCommitToRepositoryID sets the "RepoCommitToRepository" edge to the Repository entity by ID if the given value is not nil.
func (rcuo *RepoCommitUpdateOne) SetNillableRepoCommitToRepositoryID(id *uuid.UUID) *RepoCommitUpdateOne {
	if id != nil {
		rcuo = rcuo.SetRepoCommitToRepositoryID(*id)
	}
	return rcuo
}

// SetRepoCommitToRepository sets the "RepoCommitToRepository" edge to the Repository entity.
func (rcuo *RepoCommitUpdateOne) SetRepoCommitToRepository(r *Repository) *RepoCommitUpdateOne {
	return rcuo.SetRepoCommitToRepositoryID(r.ID)
}

// Mutation returns the RepoCommitMutation object of the builder.
func (rcuo *RepoCommitUpdateOne) Mutation() *RepoCommitMutation {
	return rcuo.mutation
}

// ClearRepoCommitToRepository clears the "RepoCommitToRepository" edge to the Repository entity.
func (rcuo *RepoCommitUpdateOne) ClearRepoCommitToRepository() *RepoCommitUpdateOne {
	rcuo.mutation.ClearRepoCommitToRepository()
	return rcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *RepoCommitUpdateOne) Select(field string, fields ...string) *RepoCommitUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated RepoCommit entity.
func (rcuo *RepoCommitUpdateOne) Save(ctx context.Context) (*RepoCommit, error) {
	var (
		err  error
		node *RepoCommit
	)
	if len(rcuo.hooks) == 0 {
		node, err = rcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rcuo.mutation = mutation
			node, err = rcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rcuo.hooks) - 1; i >= 0; i-- {
			if rcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *RepoCommitUpdateOne) SaveX(ctx context.Context) *RepoCommit {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *RepoCommitUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *RepoCommitUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rcuo *RepoCommitUpdateOne) sqlSave(ctx context.Context) (_node *RepoCommit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repocommit.Table,
			Columns: repocommit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: repocommit.FieldID,
			},
		},
	}
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RepoCommit.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repocommit.FieldID)
		for _, f := range fields {
			if !repocommit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repocommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repocommit.FieldRevision,
		})
	}
	if value, ok := rcuo.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: repocommit.FieldRevision,
		})
	}
	if value, ok := rcuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldHash,
		})
	}
	if value, ok := rcuo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: repocommit.FieldAuthor,
		})
	}
	if value, ok := rcuo.mutation.Committer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: repocommit.FieldCommitter,
		})
	}
	if value, ok := rcuo.mutation.PgpSignature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldPgpSignature,
		})
	}
	if value, ok := rcuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldMessage,
		})
	}
	if value, ok := rcuo.mutation.TreeHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repocommit.FieldTreeHash,
		})
	}
	if value, ok := rcuo.mutation.ParentHashes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: repocommit.FieldParentHashes,
		})
	}
	if rcuo.mutation.RepoCommitToRepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repocommit.RepoCommitToRepositoryTable,
			Columns: []string{repocommit.RepoCommitToRepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rcuo.mutation.RepoCommitToRepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repocommit.RepoCommitToRepositoryTable,
			Columns: []string{repocommit.RepoCommitToRepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RepoCommit{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repocommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
