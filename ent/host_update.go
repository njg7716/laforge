// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// HostUpdate is the builder for updating Host entities.
type HostUpdate struct {
	config
	hooks    []Hook
	mutation *HostMutation
}

// Where adds a new predicate for the builder.
func (hu *HostUpdate) Where(ps ...predicate.Host) *HostUpdate {
	hu.mutation.predicates = append(hu.mutation.predicates, ps...)
	return hu
}

// SetHostname sets the hostname field.
func (hu *HostUpdate) SetHostname(s string) *HostUpdate {
	hu.mutation.SetHostname(s)
	return hu
}

// SetDescription sets the description field.
func (hu *HostUpdate) SetDescription(s string) *HostUpdate {
	hu.mutation.SetDescription(s)
	return hu
}

// SetString sets the string field.
func (hu *HostUpdate) SetString(s string) *HostUpdate {
	hu.mutation.SetString(s)
	return hu
}

// SetLastOctet sets the last_octet field.
func (hu *HostUpdate) SetLastOctet(i int) *HostUpdate {
	hu.mutation.ResetLastOctet()
	hu.mutation.SetLastOctet(i)
	return hu
}

// AddLastOctet adds i to last_octet.
func (hu *HostUpdate) AddLastOctet(i int) *HostUpdate {
	hu.mutation.AddLastOctet(i)
	return hu
}

// SetAllowMACChanges sets the allow_mac_changes field.
func (hu *HostUpdate) SetAllowMACChanges(b bool) *HostUpdate {
	hu.mutation.SetAllowMACChanges(b)
	return hu
}

// SetExposedTCPPorts sets the exposed_tcp_ports field.
func (hu *HostUpdate) SetExposedTCPPorts(s []string) *HostUpdate {
	hu.mutation.SetExposedTCPPorts(s)
	return hu
}

// SetExposedUDPPorts sets the exposed_udp_ports field.
func (hu *HostUpdate) SetExposedUDPPorts(s []string) *HostUpdate {
	hu.mutation.SetExposedUDPPorts(s)
	return hu
}

// SetOverridePassword sets the override_password field.
func (hu *HostUpdate) SetOverridePassword(s string) *HostUpdate {
	hu.mutation.SetOverridePassword(s)
	return hu
}

// SetVars sets the vars field.
func (hu *HostUpdate) SetVars(s []string) *HostUpdate {
	hu.mutation.SetVars(s)
	return hu
}

// SetUserGroups sets the user_groups field.
func (hu *HostUpdate) SetUserGroups(s []string) *HostUpdate {
	hu.mutation.SetUserGroups(s)
	return hu
}

// SetDependsOn sets the depends_on field.
func (hu *HostUpdate) SetDependsOn(s []string) *HostUpdate {
	hu.mutation.SetDependsOn(s)
	return hu
}

// SetScripts sets the scripts field.
func (hu *HostUpdate) SetScripts(s []string) *HostUpdate {
	hu.mutation.SetScripts(s)
	return hu
}

// SetCommands sets the commands field.
func (hu *HostUpdate) SetCommands(s []string) *HostUpdate {
	hu.mutation.SetCommands(s)
	return hu
}

// SetRemoteFiles sets the remote_files field.
func (hu *HostUpdate) SetRemoteFiles(s []string) *HostUpdate {
	hu.mutation.SetRemoteFiles(s)
	return hu
}

// SetDNSRecords sets the dns_records field.
func (hu *HostUpdate) SetDNSRecords(s []string) *HostUpdate {
	hu.mutation.SetDNSRecords(s)
	return hu
}

// AddDiskIDs adds the disk edge to Disk by ids.
func (hu *HostUpdate) AddDiskIDs(ids ...int) *HostUpdate {
	hu.mutation.AddDiskIDs(ids...)
	return hu
}

// AddDisk adds the disk edges to Disk.
func (hu *HostUpdate) AddDisk(d ...*Disk) *HostUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return hu.AddDiskIDs(ids...)
}

// AddMaintainerIDs adds the maintainer edge to User by ids.
func (hu *HostUpdate) AddMaintainerIDs(ids ...int) *HostUpdate {
	hu.mutation.AddMaintainerIDs(ids...)
	return hu
}

// AddMaintainer adds the maintainer edges to User.
func (hu *HostUpdate) AddMaintainer(u ...*User) *HostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hu.AddMaintainerIDs(ids...)
}

// AddTagIDs adds the tag edge to Tag by ids.
func (hu *HostUpdate) AddTagIDs(ids ...int) *HostUpdate {
	hu.mutation.AddTagIDs(ids...)
	return hu
}

// AddTag adds the tag edges to Tag.
func (hu *HostUpdate) AddTag(t ...*Tag) *HostUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return hu.AddTagIDs(ids...)
}

// Mutation returns the HostMutation object of the builder.
func (hu *HostUpdate) Mutation() *HostMutation {
	return hu.mutation
}

// ClearDisk clears all "disk" edges to type Disk.
func (hu *HostUpdate) ClearDisk() *HostUpdate {
	hu.mutation.ClearDisk()
	return hu
}

// RemoveDiskIDs removes the disk edge to Disk by ids.
func (hu *HostUpdate) RemoveDiskIDs(ids ...int) *HostUpdate {
	hu.mutation.RemoveDiskIDs(ids...)
	return hu
}

// RemoveDisk removes disk edges to Disk.
func (hu *HostUpdate) RemoveDisk(d ...*Disk) *HostUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return hu.RemoveDiskIDs(ids...)
}

// ClearMaintainer clears all "maintainer" edges to type User.
func (hu *HostUpdate) ClearMaintainer() *HostUpdate {
	hu.mutation.ClearMaintainer()
	return hu
}

// RemoveMaintainerIDs removes the maintainer edge to User by ids.
func (hu *HostUpdate) RemoveMaintainerIDs(ids ...int) *HostUpdate {
	hu.mutation.RemoveMaintainerIDs(ids...)
	return hu
}

// RemoveMaintainer removes maintainer edges to User.
func (hu *HostUpdate) RemoveMaintainer(u ...*User) *HostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hu.RemoveMaintainerIDs(ids...)
}

// ClearTag clears all "tag" edges to type Tag.
func (hu *HostUpdate) ClearTag() *HostUpdate {
	hu.mutation.ClearTag()
	return hu
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (hu *HostUpdate) RemoveTagIDs(ids ...int) *HostUpdate {
	hu.mutation.RemoveTagIDs(ids...)
	return hu
}

// RemoveTag removes tag edges to Tag.
func (hu *HostUpdate) RemoveTag(t ...*Tag) *HostUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return hu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HostUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HostUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HostUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   host.Table,
			Columns: host.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: host.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Hostname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldHostname,
		})
	}
	if value, ok := hu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldDescription,
		})
	}
	if value, ok := hu.mutation.String(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldString,
		})
	}
	if value, ok := hu.mutation.LastOctet(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: host.FieldLastOctet,
		})
	}
	if value, ok := hu.mutation.AddedLastOctet(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: host.FieldLastOctet,
		})
	}
	if value, ok := hu.mutation.AllowMACChanges(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: host.FieldAllowMACChanges,
		})
	}
	if value, ok := hu.mutation.ExposedTCPPorts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldExposedTCPPorts,
		})
	}
	if value, ok := hu.mutation.ExposedUDPPorts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldExposedUDPPorts,
		})
	}
	if value, ok := hu.mutation.OverridePassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldOverridePassword,
		})
	}
	if value, ok := hu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldVars,
		})
	}
	if value, ok := hu.mutation.UserGroups(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldUserGroups,
		})
	}
	if value, ok := hu.mutation.DependsOn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldDependsOn,
		})
	}
	if value, ok := hu.mutation.Scripts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldScripts,
		})
	}
	if value, ok := hu.mutation.Commands(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldCommands,
		})
	}
	if value, ok := hu.mutation.RemoteFiles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldRemoteFiles,
		})
	}
	if value, ok := hu.mutation.DNSRecords(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldDNSRecords,
		})
	}
	if hu.mutation.DiskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedDiskIDs(); len(nodes) > 0 && !hu.mutation.DiskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.DiskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.MaintainerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
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
	if nodes := hu.mutation.RemovedMaintainerIDs(); len(nodes) > 0 && !hu.mutation.MaintainerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
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
	if nodes := hu.mutation.MaintainerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
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
	if hu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
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
	if nodes := hu.mutation.RemovedTagIDs(); len(nodes) > 0 && !hu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
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
	if nodes := hu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HostUpdateOne is the builder for updating a single Host entity.
type HostUpdateOne struct {
	config
	hooks    []Hook
	mutation *HostMutation
}

// SetHostname sets the hostname field.
func (huo *HostUpdateOne) SetHostname(s string) *HostUpdateOne {
	huo.mutation.SetHostname(s)
	return huo
}

// SetDescription sets the description field.
func (huo *HostUpdateOne) SetDescription(s string) *HostUpdateOne {
	huo.mutation.SetDescription(s)
	return huo
}

// SetString sets the string field.
func (huo *HostUpdateOne) SetString(s string) *HostUpdateOne {
	huo.mutation.SetString(s)
	return huo
}

// SetLastOctet sets the last_octet field.
func (huo *HostUpdateOne) SetLastOctet(i int) *HostUpdateOne {
	huo.mutation.ResetLastOctet()
	huo.mutation.SetLastOctet(i)
	return huo
}

// AddLastOctet adds i to last_octet.
func (huo *HostUpdateOne) AddLastOctet(i int) *HostUpdateOne {
	huo.mutation.AddLastOctet(i)
	return huo
}

// SetAllowMACChanges sets the allow_mac_changes field.
func (huo *HostUpdateOne) SetAllowMACChanges(b bool) *HostUpdateOne {
	huo.mutation.SetAllowMACChanges(b)
	return huo
}

// SetExposedTCPPorts sets the exposed_tcp_ports field.
func (huo *HostUpdateOne) SetExposedTCPPorts(s []string) *HostUpdateOne {
	huo.mutation.SetExposedTCPPorts(s)
	return huo
}

// SetExposedUDPPorts sets the exposed_udp_ports field.
func (huo *HostUpdateOne) SetExposedUDPPorts(s []string) *HostUpdateOne {
	huo.mutation.SetExposedUDPPorts(s)
	return huo
}

// SetOverridePassword sets the override_password field.
func (huo *HostUpdateOne) SetOverridePassword(s string) *HostUpdateOne {
	huo.mutation.SetOverridePassword(s)
	return huo
}

// SetVars sets the vars field.
func (huo *HostUpdateOne) SetVars(s []string) *HostUpdateOne {
	huo.mutation.SetVars(s)
	return huo
}

// SetUserGroups sets the user_groups field.
func (huo *HostUpdateOne) SetUserGroups(s []string) *HostUpdateOne {
	huo.mutation.SetUserGroups(s)
	return huo
}

// SetDependsOn sets the depends_on field.
func (huo *HostUpdateOne) SetDependsOn(s []string) *HostUpdateOne {
	huo.mutation.SetDependsOn(s)
	return huo
}

// SetScripts sets the scripts field.
func (huo *HostUpdateOne) SetScripts(s []string) *HostUpdateOne {
	huo.mutation.SetScripts(s)
	return huo
}

// SetCommands sets the commands field.
func (huo *HostUpdateOne) SetCommands(s []string) *HostUpdateOne {
	huo.mutation.SetCommands(s)
	return huo
}

// SetRemoteFiles sets the remote_files field.
func (huo *HostUpdateOne) SetRemoteFiles(s []string) *HostUpdateOne {
	huo.mutation.SetRemoteFiles(s)
	return huo
}

// SetDNSRecords sets the dns_records field.
func (huo *HostUpdateOne) SetDNSRecords(s []string) *HostUpdateOne {
	huo.mutation.SetDNSRecords(s)
	return huo
}

// AddDiskIDs adds the disk edge to Disk by ids.
func (huo *HostUpdateOne) AddDiskIDs(ids ...int) *HostUpdateOne {
	huo.mutation.AddDiskIDs(ids...)
	return huo
}

// AddDisk adds the disk edges to Disk.
func (huo *HostUpdateOne) AddDisk(d ...*Disk) *HostUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return huo.AddDiskIDs(ids...)
}

// AddMaintainerIDs adds the maintainer edge to User by ids.
func (huo *HostUpdateOne) AddMaintainerIDs(ids ...int) *HostUpdateOne {
	huo.mutation.AddMaintainerIDs(ids...)
	return huo
}

// AddMaintainer adds the maintainer edges to User.
func (huo *HostUpdateOne) AddMaintainer(u ...*User) *HostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return huo.AddMaintainerIDs(ids...)
}

// AddTagIDs adds the tag edge to Tag by ids.
func (huo *HostUpdateOne) AddTagIDs(ids ...int) *HostUpdateOne {
	huo.mutation.AddTagIDs(ids...)
	return huo
}

// AddTag adds the tag edges to Tag.
func (huo *HostUpdateOne) AddTag(t ...*Tag) *HostUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return huo.AddTagIDs(ids...)
}

// Mutation returns the HostMutation object of the builder.
func (huo *HostUpdateOne) Mutation() *HostMutation {
	return huo.mutation
}

// ClearDisk clears all "disk" edges to type Disk.
func (huo *HostUpdateOne) ClearDisk() *HostUpdateOne {
	huo.mutation.ClearDisk()
	return huo
}

// RemoveDiskIDs removes the disk edge to Disk by ids.
func (huo *HostUpdateOne) RemoveDiskIDs(ids ...int) *HostUpdateOne {
	huo.mutation.RemoveDiskIDs(ids...)
	return huo
}

// RemoveDisk removes disk edges to Disk.
func (huo *HostUpdateOne) RemoveDisk(d ...*Disk) *HostUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return huo.RemoveDiskIDs(ids...)
}

// ClearMaintainer clears all "maintainer" edges to type User.
func (huo *HostUpdateOne) ClearMaintainer() *HostUpdateOne {
	huo.mutation.ClearMaintainer()
	return huo
}

// RemoveMaintainerIDs removes the maintainer edge to User by ids.
func (huo *HostUpdateOne) RemoveMaintainerIDs(ids ...int) *HostUpdateOne {
	huo.mutation.RemoveMaintainerIDs(ids...)
	return huo
}

// RemoveMaintainer removes maintainer edges to User.
func (huo *HostUpdateOne) RemoveMaintainer(u ...*User) *HostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return huo.RemoveMaintainerIDs(ids...)
}

// ClearTag clears all "tag" edges to type Tag.
func (huo *HostUpdateOne) ClearTag() *HostUpdateOne {
	huo.mutation.ClearTag()
	return huo
}

// RemoveTagIDs removes the tag edge to Tag by ids.
func (huo *HostUpdateOne) RemoveTagIDs(ids ...int) *HostUpdateOne {
	huo.mutation.RemoveTagIDs(ids...)
	return huo
}

// RemoveTag removes tag edges to Tag.
func (huo *HostUpdateOne) RemoveTag(t ...*Tag) *HostUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return huo.RemoveTagIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (huo *HostUpdateOne) Save(ctx context.Context) (*Host, error) {
	var (
		err  error
		node *Host
	)
	if len(huo.hooks) == 0 {
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HostUpdateOne) SaveX(ctx context.Context) *Host {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HostUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HostUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HostUpdateOne) sqlSave(ctx context.Context) (_node *Host, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   host.Table,
			Columns: host.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: host.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Host.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := huo.mutation.Hostname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldHostname,
		})
	}
	if value, ok := huo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldDescription,
		})
	}
	if value, ok := huo.mutation.String(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldString,
		})
	}
	if value, ok := huo.mutation.LastOctet(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: host.FieldLastOctet,
		})
	}
	if value, ok := huo.mutation.AddedLastOctet(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: host.FieldLastOctet,
		})
	}
	if value, ok := huo.mutation.AllowMACChanges(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: host.FieldAllowMACChanges,
		})
	}
	if value, ok := huo.mutation.ExposedTCPPorts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldExposedTCPPorts,
		})
	}
	if value, ok := huo.mutation.ExposedUDPPorts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldExposedUDPPorts,
		})
	}
	if value, ok := huo.mutation.OverridePassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldOverridePassword,
		})
	}
	if value, ok := huo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldVars,
		})
	}
	if value, ok := huo.mutation.UserGroups(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldUserGroups,
		})
	}
	if value, ok := huo.mutation.DependsOn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldDependsOn,
		})
	}
	if value, ok := huo.mutation.Scripts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldScripts,
		})
	}
	if value, ok := huo.mutation.Commands(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldCommands,
		})
	}
	if value, ok := huo.mutation.RemoteFiles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldRemoteFiles,
		})
	}
	if value, ok := huo.mutation.DNSRecords(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldDNSRecords,
		})
	}
	if huo.mutation.DiskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedDiskIDs(); len(nodes) > 0 && !huo.mutation.DiskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.DiskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.MaintainerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
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
	if nodes := huo.mutation.RemovedMaintainerIDs(); len(nodes) > 0 && !huo.mutation.MaintainerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
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
	if nodes := huo.mutation.MaintainerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
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
	if huo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
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
	if nodes := huo.mutation.RemovedTagIDs(); len(nodes) > 0 && !huo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
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
	if nodes := huo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
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
	_node = &Host{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
