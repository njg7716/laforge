// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/buildcommit"
	"github.com/gen0cide/laforge/ent/plandiff"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/servertask"
	"github.com/google/uuid"
)

// BuildCommitQuery is the builder for querying BuildCommit entities.
type BuildCommitQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BuildCommit
	// eager-loading edges.
	withBuildCommitToBuild      *BuildQuery
	withBuildCommitToServerTask *ServerTaskQuery
	withBuildCommitToPlanDiffs  *PlanDiffQuery
	withFKs                     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BuildCommitQuery builder.
func (bcq *BuildCommitQuery) Where(ps ...predicate.BuildCommit) *BuildCommitQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit adds a limit step to the query.
func (bcq *BuildCommitQuery) Limit(limit int) *BuildCommitQuery {
	bcq.limit = &limit
	return bcq
}

// Offset adds an offset step to the query.
func (bcq *BuildCommitQuery) Offset(offset int) *BuildCommitQuery {
	bcq.offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BuildCommitQuery) Unique(unique bool) *BuildCommitQuery {
	bcq.unique = &unique
	return bcq
}

// Order adds an order step to the query.
func (bcq *BuildCommitQuery) Order(o ...OrderFunc) *BuildCommitQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// QueryBuildCommitToBuild chains the current query on the "BuildCommitToBuild" edge.
func (bcq *BuildCommitQuery) QueryBuildCommitToBuild() *BuildQuery {
	query := &BuildQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(buildcommit.Table, buildcommit.FieldID, selector),
			sqlgraph.To(build.Table, build.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, buildcommit.BuildCommitToBuildTable, buildcommit.BuildCommitToBuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBuildCommitToServerTask chains the current query on the "BuildCommitToServerTask" edge.
func (bcq *BuildCommitQuery) QueryBuildCommitToServerTask() *ServerTaskQuery {
	query := &ServerTaskQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(buildcommit.Table, buildcommit.FieldID, selector),
			sqlgraph.To(servertask.Table, servertask.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, buildcommit.BuildCommitToServerTaskTable, buildcommit.BuildCommitToServerTaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBuildCommitToPlanDiffs chains the current query on the "BuildCommitToPlanDiffs" edge.
func (bcq *BuildCommitQuery) QueryBuildCommitToPlanDiffs() *PlanDiffQuery {
	query := &PlanDiffQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(buildcommit.Table, buildcommit.FieldID, selector),
			sqlgraph.To(plandiff.Table, plandiff.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, buildcommit.BuildCommitToPlanDiffsTable, buildcommit.BuildCommitToPlanDiffsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BuildCommit entity from the query.
// Returns a *NotFoundError when no BuildCommit was found.
func (bcq *BuildCommitQuery) First(ctx context.Context) (*BuildCommit, error) {
	nodes, err := bcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{buildcommit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BuildCommitQuery) FirstX(ctx context.Context) *BuildCommit {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BuildCommit ID from the query.
// Returns a *NotFoundError when no BuildCommit ID was found.
func (bcq *BuildCommitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{buildcommit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BuildCommitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BuildCommit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BuildCommit entity is not found.
// Returns a *NotFoundError when no BuildCommit entities are found.
func (bcq *BuildCommitQuery) Only(ctx context.Context) (*BuildCommit, error) {
	nodes, err := bcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{buildcommit.Label}
	default:
		return nil, &NotSingularError{buildcommit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BuildCommitQuery) OnlyX(ctx context.Context) *BuildCommit {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BuildCommit ID in the query.
// Returns a *NotSingularError when exactly one BuildCommit ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BuildCommitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = &NotSingularError{buildcommit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BuildCommitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BuildCommits.
func (bcq *BuildCommitQuery) All(ctx context.Context) ([]*BuildCommit, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BuildCommitQuery) AllX(ctx context.Context) []*BuildCommit {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BuildCommit IDs.
func (bcq *BuildCommitQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := bcq.Select(buildcommit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BuildCommitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BuildCommitQuery) Count(ctx context.Context) (int, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BuildCommitQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BuildCommitQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BuildCommitQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BuildCommitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BuildCommitQuery) Clone() *BuildCommitQuery {
	if bcq == nil {
		return nil
	}
	return &BuildCommitQuery{
		config:                      bcq.config,
		limit:                       bcq.limit,
		offset:                      bcq.offset,
		order:                       append([]OrderFunc{}, bcq.order...),
		predicates:                  append([]predicate.BuildCommit{}, bcq.predicates...),
		withBuildCommitToBuild:      bcq.withBuildCommitToBuild.Clone(),
		withBuildCommitToServerTask: bcq.withBuildCommitToServerTask.Clone(),
		withBuildCommitToPlanDiffs:  bcq.withBuildCommitToPlanDiffs.Clone(),
		// clone intermediate query.
		sql:  bcq.sql.Clone(),
		path: bcq.path,
	}
}

// WithBuildCommitToBuild tells the query-builder to eager-load the nodes that are connected to
// the "BuildCommitToBuild" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BuildCommitQuery) WithBuildCommitToBuild(opts ...func(*BuildQuery)) *BuildCommitQuery {
	query := &BuildQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBuildCommitToBuild = query
	return bcq
}

// WithBuildCommitToServerTask tells the query-builder to eager-load the nodes that are connected to
// the "BuildCommitToServerTask" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BuildCommitQuery) WithBuildCommitToServerTask(opts ...func(*ServerTaskQuery)) *BuildCommitQuery {
	query := &ServerTaskQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBuildCommitToServerTask = query
	return bcq
}

// WithBuildCommitToPlanDiffs tells the query-builder to eager-load the nodes that are connected to
// the "BuildCommitToPlanDiffs" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BuildCommitQuery) WithBuildCommitToPlanDiffs(opts ...func(*PlanDiffQuery)) *BuildCommitQuery {
	query := &PlanDiffQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBuildCommitToPlanDiffs = query
	return bcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type buildcommit.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BuildCommit.Query().
//		GroupBy(buildcommit.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bcq *BuildCommitQuery) GroupBy(field string, fields ...string) *BuildCommitGroupBy {
	group := &BuildCommitGroupBy{config: bcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type buildcommit.Type `json:"type,omitempty"`
//	}
//
//	client.BuildCommit.Query().
//		Select(buildcommit.FieldType).
//		Scan(ctx, &v)
//
func (bcq *BuildCommitQuery) Select(fields ...string) *BuildCommitSelect {
	bcq.fields = append(bcq.fields, fields...)
	return &BuildCommitSelect{BuildCommitQuery: bcq}
}

func (bcq *BuildCommitQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcq.fields {
		if !buildcommit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bcq.path != nil {
		prev, err := bcq.path(ctx)
		if err != nil {
			return err
		}
		bcq.sql = prev
	}
	return nil
}

func (bcq *BuildCommitQuery) sqlAll(ctx context.Context) ([]*BuildCommit, error) {
	var (
		nodes       = []*BuildCommit{}
		withFKs     = bcq.withFKs
		_spec       = bcq.querySpec()
		loadedTypes = [3]bool{
			bcq.withBuildCommitToBuild != nil,
			bcq.withBuildCommitToServerTask != nil,
			bcq.withBuildCommitToPlanDiffs != nil,
		}
	)
	if bcq.withBuildCommitToBuild != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, buildcommit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BuildCommit{config: bcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, bcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bcq.withBuildCommitToBuild; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*BuildCommit)
		for i := range nodes {
			if nodes[i].build_commit_build_commit_to_build == nil {
				continue
			}
			fk := *nodes[i].build_commit_build_commit_to_build
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(build.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "build_commit_build_commit_to_build" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BuildCommitToBuild = n
			}
		}
	}

	if query := bcq.withBuildCommitToServerTask; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*BuildCommit)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BuildCommitToServerTask = []*ServerTask{}
		}
		query.withFKs = true
		query.Where(predicate.ServerTask(func(s *sql.Selector) {
			s.Where(sql.InValues(buildcommit.BuildCommitToServerTaskColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.server_task_server_task_to_build_commit
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "server_task_server_task_to_build_commit" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "server_task_server_task_to_build_commit" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BuildCommitToServerTask = append(node.Edges.BuildCommitToServerTask, n)
		}
	}

	if query := bcq.withBuildCommitToPlanDiffs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*BuildCommit)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BuildCommitToPlanDiffs = []*PlanDiff{}
		}
		query.withFKs = true
		query.Where(predicate.PlanDiff(func(s *sql.Selector) {
			s.Where(sql.InValues(buildcommit.BuildCommitToPlanDiffsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.plan_diff_plan_diff_to_build_commit
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "plan_diff_plan_diff_to_build_commit" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "plan_diff_plan_diff_to_build_commit" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.BuildCommitToPlanDiffs = append(node.Edges.BuildCommitToPlanDiffs, n)
		}
	}

	return nodes, nil
}

func (bcq *BuildCommitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BuildCommitQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bcq *BuildCommitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   buildcommit.Table,
			Columns: buildcommit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: buildcommit.FieldID,
			},
		},
		From:   bcq.sql,
		Unique: true,
	}
	if unique := bcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, buildcommit.FieldID)
		for i := range fields {
			if fields[i] != buildcommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcq *BuildCommitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(buildcommit.Table)
	columns := bcq.fields
	if len(columns) == 0 {
		columns = buildcommit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcq.sql != nil {
		selector = bcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range bcq.predicates {
		p(selector)
	}
	for _, p := range bcq.order {
		p(selector)
	}
	if offset := bcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BuildCommitGroupBy is the group-by builder for BuildCommit entities.
type BuildCommitGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BuildCommitGroupBy) Aggregate(fns ...AggregateFunc) *BuildCommitGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcgb *BuildCommitGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bcgb.path(ctx)
	if err != nil {
		return err
	}
	bcgb.sql = query
	return bcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bcgb.fields) > 1 {
		return nil, errors.New("ent: BuildCommitGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) StringsX(ctx context.Context) []string {
	v, err := bcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) StringX(ctx context.Context) string {
	v, err := bcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bcgb.fields) > 1 {
		return nil, errors.New("ent: BuildCommitGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) IntsX(ctx context.Context) []int {
	v, err := bcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) IntX(ctx context.Context) int {
	v, err := bcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bcgb.fields) > 1 {
		return nil, errors.New("ent: BuildCommitGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bcgb.fields) > 1 {
		return nil, errors.New("ent: BuildCommitGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bcgb *BuildCommitGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bcgb *BuildCommitGroupBy) BoolX(ctx context.Context) bool {
	v, err := bcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bcgb *BuildCommitGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bcgb.fields {
		if !buildcommit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcgb *BuildCommitGroupBy) sqlQuery() *sql.Selector {
	selector := bcgb.sql.Select()
	aggregation := make([]string, 0, len(bcgb.fns))
	for _, fn := range bcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcgb.fields)+len(bcgb.fns))
		for _, f := range bcgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcgb.fields...)...)
}

// BuildCommitSelect is the builder for selecting fields of BuildCommit entities.
type BuildCommitSelect struct {
	*BuildCommitQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BuildCommitSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	bcs.sql = bcs.BuildCommitQuery.sqlQuery(ctx)
	return bcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bcs *BuildCommitSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bcs.fields) > 1 {
		return nil, errors.New("ent: BuildCommitSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bcs *BuildCommitSelect) StringsX(ctx context.Context) []string {
	v, err := bcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bcs *BuildCommitSelect) StringX(ctx context.Context) string {
	v, err := bcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bcs.fields) > 1 {
		return nil, errors.New("ent: BuildCommitSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bcs *BuildCommitSelect) IntsX(ctx context.Context) []int {
	v, err := bcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bcs *BuildCommitSelect) IntX(ctx context.Context) int {
	v, err := bcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bcs.fields) > 1 {
		return nil, errors.New("ent: BuildCommitSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bcs *BuildCommitSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bcs *BuildCommitSelect) Float64X(ctx context.Context) float64 {
	v, err := bcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bcs.fields) > 1 {
		return nil, errors.New("ent: BuildCommitSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bcs *BuildCommitSelect) BoolsX(ctx context.Context) []bool {
	v, err := bcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bcs *BuildCommitSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{buildcommit.Label}
	default:
		err = fmt.Errorf("ent: BuildCommitSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bcs *BuildCommitSelect) BoolX(ctx context.Context) bool {
	v, err := bcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bcs *BuildCommitSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bcs.sql.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
