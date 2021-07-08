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
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/provisioningstep"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/google/uuid"
)

// StatusCreate is the builder for creating a Status entity.
type StatusCreate struct {
	config
	mutation *StatusMutation
	hooks    []Hook
}

// SetState sets the "state" field.
func (sc *StatusCreate) SetState(s status.State) *StatusCreate {
	sc.mutation.SetState(s)
	return sc
}

// SetStatusFor sets the "status_for" field.
func (sc *StatusCreate) SetStatusFor(sf status.StatusFor) *StatusCreate {
	sc.mutation.SetStatusFor(sf)
	return sc
}

// SetStartedAt sets the "started_at" field.
func (sc *StatusCreate) SetStartedAt(t time.Time) *StatusCreate {
	sc.mutation.SetStartedAt(t)
	return sc
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (sc *StatusCreate) SetNillableStartedAt(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetStartedAt(*t)
	}
	return sc
}

// SetEndedAt sets the "ended_at" field.
func (sc *StatusCreate) SetEndedAt(t time.Time) *StatusCreate {
	sc.mutation.SetEndedAt(t)
	return sc
}

// SetNillableEndedAt sets the "ended_at" field if the given value is not nil.
func (sc *StatusCreate) SetNillableEndedAt(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetEndedAt(*t)
	}
	return sc
}

// SetFailed sets the "failed" field.
func (sc *StatusCreate) SetFailed(b bool) *StatusCreate {
	sc.mutation.SetFailed(b)
	return sc
}

// SetNillableFailed sets the "failed" field if the given value is not nil.
func (sc *StatusCreate) SetNillableFailed(b *bool) *StatusCreate {
	if b != nil {
		sc.SetFailed(*b)
	}
	return sc
}

// SetCompleted sets the "completed" field.
func (sc *StatusCreate) SetCompleted(b bool) *StatusCreate {
	sc.mutation.SetCompleted(b)
	return sc
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (sc *StatusCreate) SetNillableCompleted(b *bool) *StatusCreate {
	if b != nil {
		sc.SetCompleted(*b)
	}
	return sc
}

// SetError sets the "error" field.
func (sc *StatusCreate) SetError(s string) *StatusCreate {
	sc.mutation.SetError(s)
	return sc
}

// SetNillableError sets the "error" field if the given value is not nil.
func (sc *StatusCreate) SetNillableError(s *string) *StatusCreate {
	if s != nil {
		sc.SetError(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StatusCreate) SetID(u uuid.UUID) *StatusCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetStatusToBuildID sets the "StatusToBuild" edge to the Build entity by ID.
func (sc *StatusCreate) SetStatusToBuildID(id uuid.UUID) *StatusCreate {
	sc.mutation.SetStatusToBuildID(id)
	return sc
}

// SetNillableStatusToBuildID sets the "StatusToBuild" edge to the Build entity by ID if the given value is not nil.
func (sc *StatusCreate) SetNillableStatusToBuildID(id *uuid.UUID) *StatusCreate {
	if id != nil {
		sc = sc.SetStatusToBuildID(*id)
	}
	return sc
}

// SetStatusToBuild sets the "StatusToBuild" edge to the Build entity.
func (sc *StatusCreate) SetStatusToBuild(b *Build) *StatusCreate {
	return sc.SetStatusToBuildID(b.ID)
}

// SetStatusToProvisionedNetworkID sets the "StatusToProvisionedNetwork" edge to the ProvisionedNetwork entity by ID.
func (sc *StatusCreate) SetStatusToProvisionedNetworkID(id uuid.UUID) *StatusCreate {
	sc.mutation.SetStatusToProvisionedNetworkID(id)
	return sc
}

// SetNillableStatusToProvisionedNetworkID sets the "StatusToProvisionedNetwork" edge to the ProvisionedNetwork entity by ID if the given value is not nil.
func (sc *StatusCreate) SetNillableStatusToProvisionedNetworkID(id *uuid.UUID) *StatusCreate {
	if id != nil {
		sc = sc.SetStatusToProvisionedNetworkID(*id)
	}
	return sc
}

// SetStatusToProvisionedNetwork sets the "StatusToProvisionedNetwork" edge to the ProvisionedNetwork entity.
func (sc *StatusCreate) SetStatusToProvisionedNetwork(p *ProvisionedNetwork) *StatusCreate {
	return sc.SetStatusToProvisionedNetworkID(p.ID)
}

// SetStatusToProvisionedHostID sets the "StatusToProvisionedHost" edge to the ProvisionedHost entity by ID.
func (sc *StatusCreate) SetStatusToProvisionedHostID(id uuid.UUID) *StatusCreate {
	sc.mutation.SetStatusToProvisionedHostID(id)
	return sc
}

// SetNillableStatusToProvisionedHostID sets the "StatusToProvisionedHost" edge to the ProvisionedHost entity by ID if the given value is not nil.
func (sc *StatusCreate) SetNillableStatusToProvisionedHostID(id *uuid.UUID) *StatusCreate {
	if id != nil {
		sc = sc.SetStatusToProvisionedHostID(*id)
	}
	return sc
}

// SetStatusToProvisionedHost sets the "StatusToProvisionedHost" edge to the ProvisionedHost entity.
func (sc *StatusCreate) SetStatusToProvisionedHost(p *ProvisionedHost) *StatusCreate {
	return sc.SetStatusToProvisionedHostID(p.ID)
}

// SetStatusToProvisioningStepID sets the "StatusToProvisioningStep" edge to the ProvisioningStep entity by ID.
func (sc *StatusCreate) SetStatusToProvisioningStepID(id uuid.UUID) *StatusCreate {
	sc.mutation.SetStatusToProvisioningStepID(id)
	return sc
}

// SetNillableStatusToProvisioningStepID sets the "StatusToProvisioningStep" edge to the ProvisioningStep entity by ID if the given value is not nil.
func (sc *StatusCreate) SetNillableStatusToProvisioningStepID(id *uuid.UUID) *StatusCreate {
	if id != nil {
		sc = sc.SetStatusToProvisioningStepID(*id)
	}
	return sc
}

// SetStatusToProvisioningStep sets the "StatusToProvisioningStep" edge to the ProvisioningStep entity.
func (sc *StatusCreate) SetStatusToProvisioningStep(p *ProvisioningStep) *StatusCreate {
	return sc.SetStatusToProvisioningStepID(p.ID)
}

// SetStatusToTeamID sets the "StatusToTeam" edge to the Team entity by ID.
func (sc *StatusCreate) SetStatusToTeamID(id uuid.UUID) *StatusCreate {
	sc.mutation.SetStatusToTeamID(id)
	return sc
}

// SetNillableStatusToTeamID sets the "StatusToTeam" edge to the Team entity by ID if the given value is not nil.
func (sc *StatusCreate) SetNillableStatusToTeamID(id *uuid.UUID) *StatusCreate {
	if id != nil {
		sc = sc.SetStatusToTeamID(*id)
	}
	return sc
}

// SetStatusToTeam sets the "StatusToTeam" edge to the Team entity.
func (sc *StatusCreate) SetStatusToTeam(t *Team) *StatusCreate {
	return sc.SetStatusToTeamID(t.ID)
}

// SetStatusToPlanID sets the "StatusToPlan" edge to the Plan entity by ID.
func (sc *StatusCreate) SetStatusToPlanID(id uuid.UUID) *StatusCreate {
	sc.mutation.SetStatusToPlanID(id)
	return sc
}

// SetNillableStatusToPlanID sets the "StatusToPlan" edge to the Plan entity by ID if the given value is not nil.
func (sc *StatusCreate) SetNillableStatusToPlanID(id *uuid.UUID) *StatusCreate {
	if id != nil {
		sc = sc.SetStatusToPlanID(*id)
	}
	return sc
}

// SetStatusToPlan sets the "StatusToPlan" edge to the Plan entity.
func (sc *StatusCreate) SetStatusToPlan(p *Plan) *StatusCreate {
	return sc.SetStatusToPlanID(p.ID)
}

// Mutation returns the StatusMutation object of the builder.
func (sc *StatusCreate) Mutation() *StatusMutation {
	return sc.mutation
}

// Save creates the Status in the database.
func (sc *StatusCreate) Save(ctx context.Context) (*Status, error) {
	var (
		err  error
		node *Status
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StatusCreate) SaveX(ctx context.Context) *Status {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sc *StatusCreate) defaults() {
	if _, ok := sc.mutation.Failed(); !ok {
		v := status.DefaultFailed
		sc.mutation.SetFailed(v)
	}
	if _, ok := sc.mutation.Completed(); !ok {
		v := status.DefaultCompleted
		sc.mutation.SetCompleted(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := status.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StatusCreate) check() error {
	if _, ok := sc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New("ent: missing required field \"state\"")}
	}
	if v, ok := sc.mutation.State(); ok {
		if err := status.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if _, ok := sc.mutation.StatusFor(); !ok {
		return &ValidationError{Name: "status_for", err: errors.New("ent: missing required field \"status_for\"")}
	}
	if v, ok := sc.mutation.StatusFor(); ok {
		if err := status.StatusForValidator(v); err != nil {
			return &ValidationError{Name: "status_for", err: fmt.Errorf("ent: validator failed for field \"status_for\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Failed(); !ok {
		return &ValidationError{Name: "failed", err: errors.New("ent: missing required field \"failed\"")}
	}
	if _, ok := sc.mutation.Completed(); !ok {
		return &ValidationError{Name: "completed", err: errors.New("ent: missing required field \"completed\"")}
	}
	return nil
}

func (sc *StatusCreate) sqlSave(ctx context.Context) (*Status, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (sc *StatusCreate) createSpec() (*Status, *sqlgraph.CreateSpec) {
	var (
		_node = &Status{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: status.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: status.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: status.FieldState,
		})
		_node.State = value
	}
	if value, ok := sc.mutation.StatusFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: status.FieldStatusFor,
		})
		_node.StatusFor = value
	}
	if value, ok := sc.mutation.StartedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: status.FieldStartedAt,
		})
		_node.StartedAt = value
	}
	if value, ok := sc.mutation.EndedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: status.FieldEndedAt,
		})
		_node.EndedAt = value
	}
	if value, ok := sc.mutation.Failed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: status.FieldFailed,
		})
		_node.Failed = value
	}
	if value, ok := sc.mutation.Completed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: status.FieldCompleted,
		})
		_node.Completed = value
	}
	if value, ok := sc.mutation.Error(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: status.FieldError,
		})
		_node.Error = value
	}
	if nodes := sc.mutation.StatusToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   status.StatusToBuildTable,
			Columns: []string{status.StatusToBuildColumn},
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
		_node.build_build_to_status = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   status.StatusToProvisionedNetworkTable,
			Columns: []string{status.StatusToProvisionedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioned_network_provisioned_network_to_status = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusToProvisionedHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   status.StatusToProvisionedHostTable,
			Columns: []string{status.StatusToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioned_host_provisioned_host_to_status = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusToProvisioningStepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   status.StatusToProvisioningStepTable,
			Columns: []string{status.StatusToProvisioningStepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisioningstep.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_status = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   status.StatusToTeamTable,
			Columns: []string{status.StatusToTeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.team_team_to_status = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatusToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   status.StatusToPlanTable,
			Columns: []string{status.StatusToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.plan_plan_to_status = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StatusCreateBulk is the builder for creating many Status entities in bulk.
type StatusCreateBulk struct {
	config
	builders []*StatusCreate
}

// Save creates the Status entities in the database.
func (scb *StatusCreateBulk) Save(ctx context.Context) ([]*Status, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Status, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StatusMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StatusCreateBulk) SaveX(ctx context.Context) []*Status {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
