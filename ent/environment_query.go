// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// EnvironmentQuery is the builder for querying Environment entities.
type EnvironmentQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Environment
	// eager-loading edges.
	withTag             *TagQuery
	withUser            *UserQuery
	withIncludedNetwork *IncludedNetworkQuery
	withBuild           *BuildQuery
	withNetwork         *NetworkQuery
	withHost            *HostQuery
	withCompetition     *CompetitionQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (eq *EnvironmentQuery) Where(ps ...predicate.Environment) *EnvironmentQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EnvironmentQuery) Limit(limit int) *EnvironmentQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EnvironmentQuery) Offset(offset int) *EnvironmentQuery {
	eq.offset = &offset
	return eq
}

// Order adds an order step to the query.
func (eq *EnvironmentQuery) Order(o ...OrderFunc) *EnvironmentQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryTag chains the current query on the tag edge.
func (eq *EnvironmentQuery) QueryTag() *TagQuery {
	query := &TagQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.TagTable, environment.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the user edge.
func (eq *EnvironmentQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.UserTable, environment.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIncludedNetwork chains the current query on the included_network edge.
func (eq *EnvironmentQuery) QueryIncludedNetwork() *IncludedNetworkQuery {
	query := &IncludedNetworkQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(includednetwork.Table, includednetwork.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.IncludedNetworkTable, environment.IncludedNetworkColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBuild chains the current query on the build edge.
func (eq *EnvironmentQuery) QueryBuild() *BuildQuery {
	query := &BuildQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(build.Table, build.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.BuildTable, environment.BuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNetwork chains the current query on the network edge.
func (eq *EnvironmentQuery) QueryNetwork() *NetworkQuery {
	query := &NetworkQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(network.Table, network.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.NetworkTable, environment.NetworkColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHost chains the current query on the host edge.
func (eq *EnvironmentQuery) QueryHost() *HostQuery {
	query := &HostQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.HostTable, environment.HostColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompetition chains the current query on the competition edge.
func (eq *EnvironmentQuery) QueryCompetition() *CompetitionQuery {
	query := &CompetitionQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(environment.Table, environment.FieldID, selector),
			sqlgraph.To(competition.Table, competition.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, environment.CompetitionTable, environment.CompetitionColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Environment entity in the query. Returns *NotFoundError when no environment was found.
func (eq *EnvironmentQuery) First(ctx context.Context) (*Environment, error) {
	nodes, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{environment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EnvironmentQuery) FirstX(ctx context.Context) *Environment {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Environment id in the query. Returns *NotFoundError when no id was found.
func (eq *EnvironmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{environment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EnvironmentQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Environment entity in the query, returns an error if not exactly one entity was returned.
func (eq *EnvironmentQuery) Only(ctx context.Context) (*Environment, error) {
	nodes, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{environment.Label}
	default:
		return nil, &NotSingularError{environment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EnvironmentQuery) OnlyX(ctx context.Context) *Environment {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Environment id in the query, returns an error if not exactly one id was returned.
func (eq *EnvironmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = &NotSingularError{environment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EnvironmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Environments.
func (eq *EnvironmentQuery) All(ctx context.Context) ([]*Environment, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EnvironmentQuery) AllX(ctx context.Context) []*Environment {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Environment ids.
func (eq *EnvironmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := eq.Select(environment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EnvironmentQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EnvironmentQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EnvironmentQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EnvironmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EnvironmentQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EnvironmentQuery) Clone() *EnvironmentQuery {
	if eq == nil {
		return nil
	}
	return &EnvironmentQuery{
		config:              eq.config,
		limit:               eq.limit,
		offset:              eq.offset,
		order:               append([]OrderFunc{}, eq.order...),
		unique:              append([]string{}, eq.unique...),
		predicates:          append([]predicate.Environment{}, eq.predicates...),
		withTag:             eq.withTag.Clone(),
		withUser:            eq.withUser.Clone(),
		withIncludedNetwork: eq.withIncludedNetwork.Clone(),
		withBuild:           eq.withBuild.Clone(),
		withNetwork:         eq.withNetwork.Clone(),
		withHost:            eq.withHost.Clone(),
		withCompetition:     eq.withCompetition.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

//  WithTag tells the query-builder to eager-loads the nodes that are connected to
// the "tag" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithTag(opts ...func(*TagQuery)) *EnvironmentQuery {
	query := &TagQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withTag = query
	return eq
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithUser(opts ...func(*UserQuery)) *EnvironmentQuery {
	query := &UserQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withUser = query
	return eq
}

//  WithIncludedNetwork tells the query-builder to eager-loads the nodes that are connected to
// the "included_network" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithIncludedNetwork(opts ...func(*IncludedNetworkQuery)) *EnvironmentQuery {
	query := &IncludedNetworkQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withIncludedNetwork = query
	return eq
}

//  WithBuild tells the query-builder to eager-loads the nodes that are connected to
// the "build" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithBuild(opts ...func(*BuildQuery)) *EnvironmentQuery {
	query := &BuildQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withBuild = query
	return eq
}

//  WithNetwork tells the query-builder to eager-loads the nodes that are connected to
// the "network" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithNetwork(opts ...func(*NetworkQuery)) *EnvironmentQuery {
	query := &NetworkQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withNetwork = query
	return eq
}

//  WithHost tells the query-builder to eager-loads the nodes that are connected to
// the "host" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithHost(opts ...func(*HostQuery)) *EnvironmentQuery {
	query := &HostQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withHost = query
	return eq
}

//  WithCompetition tells the query-builder to eager-loads the nodes that are connected to
// the "competition" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EnvironmentQuery) WithCompetition(opts ...func(*CompetitionQuery)) *EnvironmentQuery {
	query := &CompetitionQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withCompetition = query
	return eq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CompetitionID string `json:"competition_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Environment.Query().
//		GroupBy(environment.FieldCompetitionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (eq *EnvironmentQuery) GroupBy(field string, fields ...string) *EnvironmentGroupBy {
	group := &EnvironmentGroupBy{config: eq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CompetitionID string `json:"competition_id,omitempty"`
//	}
//
//	client.Environment.Query().
//		Select(environment.FieldCompetitionID).
//		Scan(ctx, &v)
//
func (eq *EnvironmentQuery) Select(field string, fields ...string) *EnvironmentSelect {
	selector := &EnvironmentSelect{config: eq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(), nil
	}
	return selector
}

func (eq *EnvironmentQuery) prepareQuery(ctx context.Context) error {
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EnvironmentQuery) sqlAll(ctx context.Context) ([]*Environment, error) {
	var (
		nodes       = []*Environment{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [7]bool{
			eq.withTag != nil,
			eq.withUser != nil,
			eq.withIncludedNetwork != nil,
			eq.withBuild != nil,
			eq.withNetwork != nil,
			eq.withHost != nil,
			eq.withCompetition != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, environment.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Environment{config: eq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := eq.withTag; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Tag = []*Tag{}
		}
		query.withFKs = true
		query.Where(predicate.Tag(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.TagColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_tag
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_tag" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_tag" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Tag = append(node.Edges.Tag, n)
		}
	}

	if query := eq.withUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.User = []*User{}
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.UserColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_user
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_user" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_user" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.User = append(node.Edges.User, n)
		}
	}

	if query := eq.withIncludedNetwork; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.IncludedNetwork = []*IncludedNetwork{}
		}
		query.withFKs = true
		query.Where(predicate.IncludedNetwork(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.IncludedNetworkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_included_network
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_included_network" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_included_network" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.IncludedNetwork = append(node.Edges.IncludedNetwork, n)
		}
	}

	if query := eq.withBuild; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Build = []*Build{}
		}
		query.withFKs = true
		query.Where(predicate.Build(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.BuildColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_build
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_build" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_build" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Build = append(node.Edges.Build, n)
		}
	}

	if query := eq.withNetwork; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Network = []*Network{}
		}
		query.withFKs = true
		query.Where(predicate.Network(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.NetworkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_network
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_network" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_network" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Network = append(node.Edges.Network, n)
		}
	}

	if query := eq.withHost; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Host = []*Host{}
		}
		query.withFKs = true
		query.Where(predicate.Host(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.HostColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_host
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_host" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_host" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Host = append(node.Edges.Host, n)
		}
	}

	if query := eq.withCompetition; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Environment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Competition = []*Competition{}
		}
		query.withFKs = true
		query.Where(predicate.Competition(func(s *sql.Selector) {
			s.Where(sql.InValues(environment.CompetitionColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.environment_competition
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "environment_competition" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_competition" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Competition = append(node.Edges.Competition, n)
		}
	}

	return nodes, nil
}

func (eq *EnvironmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EnvironmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (eq *EnvironmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   environment.Table,
			Columns: environment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: environment.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, environment.ValidColumn)
			}
		}
	}
	return _spec
}

func (eq *EnvironmentQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(environment.Table)
	selector := builder.Select(t1.Columns(environment.Columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(environment.Columns...)...)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector, environment.ValidColumn)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnvironmentGroupBy is the builder for group-by Environment entities.
type EnvironmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EnvironmentGroupBy) Aggregate(fns ...AggregateFunc) *EnvironmentGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scan the result into the given value.
func (egb *EnvironmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (egb *EnvironmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := egb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnvironmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (egb *EnvironmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := egb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = egb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (egb *EnvironmentGroupBy) StringX(ctx context.Context) string {
	v, err := egb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnvironmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (egb *EnvironmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := egb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = egb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (egb *EnvironmentGroupBy) IntX(ctx context.Context) int {
	v, err := egb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnvironmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (egb *EnvironmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := egb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = egb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (egb *EnvironmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := egb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnvironmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (egb *EnvironmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := egb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (egb *EnvironmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = egb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (egb *EnvironmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := egb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (egb *EnvironmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range egb.fields {
		if !environment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := egb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EnvironmentGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql
	columns := make([]string, 0, len(egb.fields)+len(egb.fns))
	columns = append(columns, egb.fields...)
	for _, fn := range egb.fns {
		columns = append(columns, fn(selector, environment.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(egb.fields...)
}

// EnvironmentSelect is the builder for select fields of Environment entities.
type EnvironmentSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (es *EnvironmentSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := es.path(ctx)
	if err != nil {
		return err
	}
	es.sql = query
	return es.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (es *EnvironmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := es.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnvironmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (es *EnvironmentSelect) StringsX(ctx context.Context) []string {
	v, err := es.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = es.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (es *EnvironmentSelect) StringX(ctx context.Context) string {
	v, err := es.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnvironmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (es *EnvironmentSelect) IntsX(ctx context.Context) []int {
	v, err := es.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = es.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (es *EnvironmentSelect) IntX(ctx context.Context) int {
	v, err := es.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnvironmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (es *EnvironmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := es.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = es.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (es *EnvironmentSelect) Float64X(ctx context.Context) float64 {
	v, err := es.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnvironmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (es *EnvironmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := es.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (es *EnvironmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = es.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{environment.Label}
	default:
		err = fmt.Errorf("ent: EnvironmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (es *EnvironmentSelect) BoolX(ctx context.Context) bool {
	v, err := es.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (es *EnvironmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range es.fields {
		if !environment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := es.sqlQuery().Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (es *EnvironmentSelect) sqlQuery() sql.Querier {
	selector := es.sql
	selector.Select(selector.Columns(es.fields...)...)
	return selector
}
