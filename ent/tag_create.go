// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/google/uuid"
)

// TagCreate is the builder for creating a Tag entity.
type TagCreate struct {
	config
	mutation *TagMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (tc *TagCreate) SetUUID(u uuid.UUID) *TagCreate {
	tc.mutation.SetUUID(u)
	return tc
}

// SetName sets the "name" field.
func (tc *TagCreate) SetName(s string) *TagCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TagCreate) SetDescription(m map[string]string) *TagCreate {
	tc.mutation.SetDescription(m)
	return tc
}

// SetID sets the "id" field.
func (tc *TagCreate) SetID(u uuid.UUID) *TagCreate {
	tc.mutation.SetID(u)
	return tc
}

// Mutation returns the TagMutation object of the builder.
func (tc *TagCreate) Mutation() *TagMutation {
	return tc.mutation
}

// Save creates the Tag in the database.
func (tc *TagCreate) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagCreate) SaveX(ctx context.Context) *Tag {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tc *TagCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := tag.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TagCreate) check() error {
	if _, ok := tc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	return nil
}

func (tc *TagCreate) sqlSave(ctx context.Context) (*Tag, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (tc *TagCreate) createSpec() (*Tag, *sqlgraph.CreateSpec) {
	var (
		_node = &Tag{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tag.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: tag.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: tag.FieldDescription,
		})
		_node.Description = value
	}
	return _node, _spec
}

// TagCreateBulk is the builder for creating many Tag entities in bulk.
type TagCreateBulk struct {
	config
	builders []*TagCreate
}

// Save creates the Tag entities in the database.
func (tcb *TagCreateBulk) Save(ctx context.Context) ([]*Tag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tag, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TagCreateBulk) SaveX(ctx context.Context) []*Tag {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
