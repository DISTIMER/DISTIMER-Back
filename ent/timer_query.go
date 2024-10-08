// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/predicate"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
)

// TimerQuery is the builder for querying Timer entities.
type TimerQuery struct {
	config
	ctx             *QueryContext
	order           []timer.OrderOption
	inters          []Interceptor
	predicates      []predicate.Timer
	withUser        *UserQuery
	withSubject     *SubjectQuery
	withSharedGroup *GroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TimerQuery builder.
func (tq *TimerQuery) Where(ps ...predicate.Timer) *TimerQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TimerQuery) Limit(limit int) *TimerQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TimerQuery) Offset(offset int) *TimerQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TimerQuery) Unique(unique bool) *TimerQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TimerQuery) Order(o ...timer.OrderOption) *TimerQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUser chains the current query on the "user" edge.
func (tq *TimerQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timer.Table, timer.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, timer.UserTable, timer.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubject chains the current query on the "subject" edge.
func (tq *TimerQuery) QuerySubject() *SubjectQuery {
	query := (&SubjectClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timer.Table, timer.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, timer.SubjectTable, timer.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySharedGroup chains the current query on the "shared_group" edge.
func (tq *TimerQuery) QuerySharedGroup() *GroupQuery {
	query := (&GroupClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timer.Table, timer.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, timer.SharedGroupTable, timer.SharedGroupPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Timer entity from the query.
// Returns a *NotFoundError when no Timer was found.
func (tq *TimerQuery) First(ctx context.Context) (*Timer, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{timer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TimerQuery) FirstX(ctx context.Context) *Timer {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Timer ID from the query.
// Returns a *NotFoundError when no Timer ID was found.
func (tq *TimerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{timer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TimerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Timer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Timer entity is found.
// Returns a *NotFoundError when no Timer entities are found.
func (tq *TimerQuery) Only(ctx context.Context) (*Timer, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{timer.Label}
	default:
		return nil, &NotSingularError{timer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TimerQuery) OnlyX(ctx context.Context) *Timer {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Timer ID in the query.
// Returns a *NotSingularError when more than one Timer ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TimerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{timer.Label}
	default:
		err = &NotSingularError{timer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TimerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Timers.
func (tq *TimerQuery) All(ctx context.Context) ([]*Timer, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Timer, *TimerQuery]()
	return withInterceptors[[]*Timer](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TimerQuery) AllX(ctx context.Context) []*Timer {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Timer IDs.
func (tq *TimerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(timer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TimerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TimerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TimerQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TimerQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TimerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TimerQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TimerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TimerQuery) Clone() *TimerQuery {
	if tq == nil {
		return nil
	}
	return &TimerQuery{
		config:          tq.config,
		ctx:             tq.ctx.Clone(),
		order:           append([]timer.OrderOption{}, tq.order...),
		inters:          append([]Interceptor{}, tq.inters...),
		predicates:      append([]predicate.Timer{}, tq.predicates...),
		withUser:        tq.withUser.Clone(),
		withSubject:     tq.withSubject.Clone(),
		withSharedGroup: tq.withSharedGroup.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimerQuery) WithUser(opts ...func(*UserQuery)) *TimerQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimerQuery) WithSubject(opts ...func(*SubjectQuery)) *TimerQuery {
	query := (&SubjectClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withSubject = query
	return tq
}

// WithSharedGroup tells the query-builder to eager-load the nodes that are connected to
// the "shared_group" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimerQuery) WithSharedGroup(opts ...func(*GroupQuery)) *TimerQuery {
	query := (&GroupClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withSharedGroup = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StartAt time.Time `json:"start_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Timer.Query().
//		GroupBy(timer.FieldStartAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TimerQuery) GroupBy(field string, fields ...string) *TimerGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TimerGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = timer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StartAt time.Time `json:"start_at,omitempty"`
//	}
//
//	client.Timer.Query().
//		Select(timer.FieldStartAt).
//		Scan(ctx, &v)
func (tq *TimerQuery) Select(fields ...string) *TimerSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TimerSelect{TimerQuery: tq}
	sbuild.label = timer.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TimerSelect configured with the given aggregations.
func (tq *TimerQuery) Aggregate(fns ...AggregateFunc) *TimerSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TimerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !timer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TimerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Timer, error) {
	var (
		nodes       = []*Timer{}
		_spec       = tq.querySpec()
		loadedTypes = [3]bool{
			tq.withUser != nil,
			tq.withSubject != nil,
			tq.withSharedGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Timer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Timer{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withUser; query != nil {
		if err := tq.loadUser(ctx, query, nodes, nil,
			func(n *Timer, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withSubject; query != nil {
		if err := tq.loadSubject(ctx, query, nodes, nil,
			func(n *Timer, e *Subject) { n.Edges.Subject = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withSharedGroup; query != nil {
		if err := tq.loadSharedGroup(ctx, query, nodes,
			func(n *Timer) { n.Edges.SharedGroup = []*Group{} },
			func(n *Timer, e *Group) { n.Edges.SharedGroup = append(n.Edges.SharedGroup, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TimerQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Timer, init func(*Timer), assign func(*Timer, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Timer)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TimerQuery) loadSubject(ctx context.Context, query *SubjectQuery, nodes []*Timer, init func(*Timer), assign func(*Timer, *Subject)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Timer)
	for i := range nodes {
		fk := nodes[i].SubjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(subject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TimerQuery) loadSharedGroup(ctx context.Context, query *GroupQuery, nodes []*Timer, init func(*Timer), assign func(*Timer, *Group)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Timer)
	nids := make(map[uuid.UUID]map[*Timer]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(timer.SharedGroupTable)
		s.Join(joinT).On(s.C(group.FieldID), joinT.C(timer.SharedGroupPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(timer.SharedGroupPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(timer.SharedGroupPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Timer]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Group](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "shared_group" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tq *TimerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TimerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(timer.Table, timer.Columns, sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timer.FieldID)
		for i := range fields {
			if fields[i] != timer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withUser != nil {
			_spec.Node.AddColumnOnce(timer.FieldUserID)
		}
		if tq.withSubject != nil {
			_spec.Node.AddColumnOnce(timer.FieldSubjectID)
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TimerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(timer.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = timer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TimerGroupBy is the group-by builder for Timer entities.
type TimerGroupBy struct {
	selector
	build *TimerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TimerGroupBy) Aggregate(fns ...AggregateFunc) *TimerGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TimerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimerQuery, *TimerGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TimerGroupBy) sqlScan(ctx context.Context, root *TimerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TimerSelect is the builder for selecting fields of Timer entities.
type TimerSelect struct {
	*TimerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TimerSelect) Aggregate(fns ...AggregateFunc) *TimerSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TimerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimerQuery, *TimerSelect](ctx, ts.TimerQuery, ts, ts.inters, v)
}

func (ts *TimerSelect) sqlScan(ctx context.Context, root *TimerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
