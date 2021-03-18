// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/dns"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// DNSUpdate is the builder for updating DNS entities.
type DNSUpdate struct {
	config
	hooks    []Hook
	mutation *DNSMutation
}

// Where adds a new predicate for the DNSUpdate builder.
func (du *DNSUpdate) Where(ps ...predicate.DNS) *DNSUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetHclID sets the "hcl_id" field.
func (du *DNSUpdate) SetHclID(s string) *DNSUpdate {
	du.mutation.SetHclID(s)
	return du
}

// SetType sets the "type" field.
func (du *DNSUpdate) SetType(s string) *DNSUpdate {
	du.mutation.SetType(s)
	return du
}

// SetRootDomain sets the "root_domain" field.
func (du *DNSUpdate) SetRootDomain(s string) *DNSUpdate {
	du.mutation.SetRootDomain(s)
	return du
}

// SetDNSServers sets the "dns_servers" field.
func (du *DNSUpdate) SetDNSServers(s []string) *DNSUpdate {
	du.mutation.SetDNSServers(s)
	return du
}

// SetNtpServers sets the "ntp_servers" field.
func (du *DNSUpdate) SetNtpServers(s []string) *DNSUpdate {
	du.mutation.SetNtpServers(s)
	return du
}

// SetConfig sets the "config" field.
func (du *DNSUpdate) SetConfig(m map[string]string) *DNSUpdate {
	du.mutation.SetConfig(m)
	return du
}

// AddDNSToTagIDs adds the "DNSToTag" edge to the Tag entity by IDs.
func (du *DNSUpdate) AddDNSToTagIDs(ids ...int) *DNSUpdate {
	du.mutation.AddDNSToTagIDs(ids...)
	return du
}

// AddDNSToTag adds the "DNSToTag" edges to the Tag entity.
func (du *DNSUpdate) AddDNSToTag(t ...*Tag) *DNSUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return du.AddDNSToTagIDs(ids...)
}

// AddDNSToEnvironmentIDs adds the "DNSToEnvironment" edge to the Environment entity by IDs.
func (du *DNSUpdate) AddDNSToEnvironmentIDs(ids ...int) *DNSUpdate {
	du.mutation.AddDNSToEnvironmentIDs(ids...)
	return du
}

// AddDNSToEnvironment adds the "DNSToEnvironment" edges to the Environment entity.
func (du *DNSUpdate) AddDNSToEnvironment(e ...*Environment) *DNSUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return du.AddDNSToEnvironmentIDs(ids...)
}

// AddDNSToCompetitionIDs adds the "DNSToCompetition" edge to the Competition entity by IDs.
func (du *DNSUpdate) AddDNSToCompetitionIDs(ids ...int) *DNSUpdate {
	du.mutation.AddDNSToCompetitionIDs(ids...)
	return du
}

// AddDNSToCompetition adds the "DNSToCompetition" edges to the Competition entity.
func (du *DNSUpdate) AddDNSToCompetition(c ...*Competition) *DNSUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.AddDNSToCompetitionIDs(ids...)
}

// Mutation returns the DNSMutation object of the builder.
func (du *DNSUpdate) Mutation() *DNSMutation {
	return du.mutation
}

// ClearDNSToTag clears all "DNSToTag" edges to the Tag entity.
func (du *DNSUpdate) ClearDNSToTag() *DNSUpdate {
	du.mutation.ClearDNSToTag()
	return du
}

// RemoveDNSToTagIDs removes the "DNSToTag" edge to Tag entities by IDs.
func (du *DNSUpdate) RemoveDNSToTagIDs(ids ...int) *DNSUpdate {
	du.mutation.RemoveDNSToTagIDs(ids...)
	return du
}

// RemoveDNSToTag removes "DNSToTag" edges to Tag entities.
func (du *DNSUpdate) RemoveDNSToTag(t ...*Tag) *DNSUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return du.RemoveDNSToTagIDs(ids...)
}

// ClearDNSToEnvironment clears all "DNSToEnvironment" edges to the Environment entity.
func (du *DNSUpdate) ClearDNSToEnvironment() *DNSUpdate {
	du.mutation.ClearDNSToEnvironment()
	return du
}

// RemoveDNSToEnvironmentIDs removes the "DNSToEnvironment" edge to Environment entities by IDs.
func (du *DNSUpdate) RemoveDNSToEnvironmentIDs(ids ...int) *DNSUpdate {
	du.mutation.RemoveDNSToEnvironmentIDs(ids...)
	return du
}

// RemoveDNSToEnvironment removes "DNSToEnvironment" edges to Environment entities.
func (du *DNSUpdate) RemoveDNSToEnvironment(e ...*Environment) *DNSUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return du.RemoveDNSToEnvironmentIDs(ids...)
}

// ClearDNSToCompetition clears all "DNSToCompetition" edges to the Competition entity.
func (du *DNSUpdate) ClearDNSToCompetition() *DNSUpdate {
	du.mutation.ClearDNSToCompetition()
	return du
}

// RemoveDNSToCompetitionIDs removes the "DNSToCompetition" edge to Competition entities by IDs.
func (du *DNSUpdate) RemoveDNSToCompetitionIDs(ids ...int) *DNSUpdate {
	du.mutation.RemoveDNSToCompetitionIDs(ids...)
	return du
}

// RemoveDNSToCompetition removes "DNSToCompetition" edges to Competition entities.
func (du *DNSUpdate) RemoveDNSToCompetition(c ...*Competition) *DNSUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return du.RemoveDNSToCompetitionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DNSUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DNSUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DNSUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DNSUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DNSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dns.Table,
			Columns: dns.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dns.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dns.FieldHclID,
		})
	}
	if value, ok := du.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dns.FieldType,
		})
	}
	if value, ok := du.mutation.RootDomain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dns.FieldRootDomain,
		})
	}
	if value, ok := du.mutation.DNSServers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dns.FieldDNSServers,
		})
	}
	if value, ok := du.mutation.NtpServers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dns.FieldNtpServers,
		})
	}
	if value, ok := du.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dns.FieldConfig,
		})
	}
	if du.mutation.DNSToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dns.DNSToTagTable,
			Columns: []string{dns.DNSToTagColumn},
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
	if nodes := du.mutation.RemovedDNSToTagIDs(); len(nodes) > 0 && !du.mutation.DNSToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dns.DNSToTagTable,
			Columns: []string{dns.DNSToTagColumn},
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
	if nodes := du.mutation.DNSToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dns.DNSToTagTable,
			Columns: []string{dns.DNSToTagColumn},
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
	if du.mutation.DNSToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToEnvironmentTable,
			Columns: dns.DNSToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDNSToEnvironmentIDs(); len(nodes) > 0 && !du.mutation.DNSToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToEnvironmentTable,
			Columns: dns.DNSToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DNSToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToEnvironmentTable,
			Columns: dns.DNSToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DNSToCompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToCompetitionTable,
			Columns: dns.DNSToCompetitionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: competition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDNSToCompetitionIDs(); len(nodes) > 0 && !du.mutation.DNSToCompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToCompetitionTable,
			Columns: dns.DNSToCompetitionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: competition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DNSToCompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToCompetitionTable,
			Columns: dns.DNSToCompetitionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: competition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dns.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DNSUpdateOne is the builder for updating a single DNS entity.
type DNSUpdateOne struct {
	config
	hooks    []Hook
	mutation *DNSMutation
}

// SetHclID sets the "hcl_id" field.
func (duo *DNSUpdateOne) SetHclID(s string) *DNSUpdateOne {
	duo.mutation.SetHclID(s)
	return duo
}

// SetType sets the "type" field.
func (duo *DNSUpdateOne) SetType(s string) *DNSUpdateOne {
	duo.mutation.SetType(s)
	return duo
}

// SetRootDomain sets the "root_domain" field.
func (duo *DNSUpdateOne) SetRootDomain(s string) *DNSUpdateOne {
	duo.mutation.SetRootDomain(s)
	return duo
}

// SetDNSServers sets the "dns_servers" field.
func (duo *DNSUpdateOne) SetDNSServers(s []string) *DNSUpdateOne {
	duo.mutation.SetDNSServers(s)
	return duo
}

// SetNtpServers sets the "ntp_servers" field.
func (duo *DNSUpdateOne) SetNtpServers(s []string) *DNSUpdateOne {
	duo.mutation.SetNtpServers(s)
	return duo
}

// SetConfig sets the "config" field.
func (duo *DNSUpdateOne) SetConfig(m map[string]string) *DNSUpdateOne {
	duo.mutation.SetConfig(m)
	return duo
}

// AddDNSToTagIDs adds the "DNSToTag" edge to the Tag entity by IDs.
func (duo *DNSUpdateOne) AddDNSToTagIDs(ids ...int) *DNSUpdateOne {
	duo.mutation.AddDNSToTagIDs(ids...)
	return duo
}

// AddDNSToTag adds the "DNSToTag" edges to the Tag entity.
func (duo *DNSUpdateOne) AddDNSToTag(t ...*Tag) *DNSUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return duo.AddDNSToTagIDs(ids...)
}

// AddDNSToEnvironmentIDs adds the "DNSToEnvironment" edge to the Environment entity by IDs.
func (duo *DNSUpdateOne) AddDNSToEnvironmentIDs(ids ...int) *DNSUpdateOne {
	duo.mutation.AddDNSToEnvironmentIDs(ids...)
	return duo
}

// AddDNSToEnvironment adds the "DNSToEnvironment" edges to the Environment entity.
func (duo *DNSUpdateOne) AddDNSToEnvironment(e ...*Environment) *DNSUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return duo.AddDNSToEnvironmentIDs(ids...)
}

// AddDNSToCompetitionIDs adds the "DNSToCompetition" edge to the Competition entity by IDs.
func (duo *DNSUpdateOne) AddDNSToCompetitionIDs(ids ...int) *DNSUpdateOne {
	duo.mutation.AddDNSToCompetitionIDs(ids...)
	return duo
}

// AddDNSToCompetition adds the "DNSToCompetition" edges to the Competition entity.
func (duo *DNSUpdateOne) AddDNSToCompetition(c ...*Competition) *DNSUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.AddDNSToCompetitionIDs(ids...)
}

// Mutation returns the DNSMutation object of the builder.
func (duo *DNSUpdateOne) Mutation() *DNSMutation {
	return duo.mutation
}

// ClearDNSToTag clears all "DNSToTag" edges to the Tag entity.
func (duo *DNSUpdateOne) ClearDNSToTag() *DNSUpdateOne {
	duo.mutation.ClearDNSToTag()
	return duo
}

// RemoveDNSToTagIDs removes the "DNSToTag" edge to Tag entities by IDs.
func (duo *DNSUpdateOne) RemoveDNSToTagIDs(ids ...int) *DNSUpdateOne {
	duo.mutation.RemoveDNSToTagIDs(ids...)
	return duo
}

// RemoveDNSToTag removes "DNSToTag" edges to Tag entities.
func (duo *DNSUpdateOne) RemoveDNSToTag(t ...*Tag) *DNSUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return duo.RemoveDNSToTagIDs(ids...)
}

// ClearDNSToEnvironment clears all "DNSToEnvironment" edges to the Environment entity.
func (duo *DNSUpdateOne) ClearDNSToEnvironment() *DNSUpdateOne {
	duo.mutation.ClearDNSToEnvironment()
	return duo
}

// RemoveDNSToEnvironmentIDs removes the "DNSToEnvironment" edge to Environment entities by IDs.
func (duo *DNSUpdateOne) RemoveDNSToEnvironmentIDs(ids ...int) *DNSUpdateOne {
	duo.mutation.RemoveDNSToEnvironmentIDs(ids...)
	return duo
}

// RemoveDNSToEnvironment removes "DNSToEnvironment" edges to Environment entities.
func (duo *DNSUpdateOne) RemoveDNSToEnvironment(e ...*Environment) *DNSUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return duo.RemoveDNSToEnvironmentIDs(ids...)
}

// ClearDNSToCompetition clears all "DNSToCompetition" edges to the Competition entity.
func (duo *DNSUpdateOne) ClearDNSToCompetition() *DNSUpdateOne {
	duo.mutation.ClearDNSToCompetition()
	return duo
}

// RemoveDNSToCompetitionIDs removes the "DNSToCompetition" edge to Competition entities by IDs.
func (duo *DNSUpdateOne) RemoveDNSToCompetitionIDs(ids ...int) *DNSUpdateOne {
	duo.mutation.RemoveDNSToCompetitionIDs(ids...)
	return duo
}

// RemoveDNSToCompetition removes "DNSToCompetition" edges to Competition entities.
func (duo *DNSUpdateOne) RemoveDNSToCompetition(c ...*Competition) *DNSUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return duo.RemoveDNSToCompetitionIDs(ids...)
}

// Save executes the query and returns the updated DNS entity.
func (duo *DNSUpdateOne) Save(ctx context.Context) (*DNS, error) {
	var (
		err  error
		node *DNS
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DNSUpdateOne) SaveX(ctx context.Context) *DNS {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DNSUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DNSUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DNSUpdateOne) sqlSave(ctx context.Context) (_node *DNS, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dns.Table,
			Columns: dns.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dns.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DNS.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dns.FieldHclID,
		})
	}
	if value, ok := duo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dns.FieldType,
		})
	}
	if value, ok := duo.mutation.RootDomain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dns.FieldRootDomain,
		})
	}
	if value, ok := duo.mutation.DNSServers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dns.FieldDNSServers,
		})
	}
	if value, ok := duo.mutation.NtpServers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dns.FieldNtpServers,
		})
	}
	if value, ok := duo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dns.FieldConfig,
		})
	}
	if duo.mutation.DNSToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dns.DNSToTagTable,
			Columns: []string{dns.DNSToTagColumn},
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
	if nodes := duo.mutation.RemovedDNSToTagIDs(); len(nodes) > 0 && !duo.mutation.DNSToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dns.DNSToTagTable,
			Columns: []string{dns.DNSToTagColumn},
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
	if nodes := duo.mutation.DNSToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dns.DNSToTagTable,
			Columns: []string{dns.DNSToTagColumn},
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
	if duo.mutation.DNSToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToEnvironmentTable,
			Columns: dns.DNSToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDNSToEnvironmentIDs(); len(nodes) > 0 && !duo.mutation.DNSToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToEnvironmentTable,
			Columns: dns.DNSToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DNSToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToEnvironmentTable,
			Columns: dns.DNSToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DNSToCompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToCompetitionTable,
			Columns: dns.DNSToCompetitionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: competition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDNSToCompetitionIDs(); len(nodes) > 0 && !duo.mutation.DNSToCompetitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToCompetitionTable,
			Columns: dns.DNSToCompetitionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: competition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DNSToCompetitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dns.DNSToCompetitionTable,
			Columns: dns.DNSToCompetitionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: competition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DNS{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dns.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
