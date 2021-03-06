// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/SrtStory/app/ent/contagious"
	"github.com/SrtStory/app/ent/predicate"
	"github.com/SrtStory/app/ent/statistic"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ContagiousQuery is the builder for querying Contagious entities.
type ContagiousQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Contagious
	// eager-loading edges.
	withStatistic *StatisticQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (cq *ContagiousQuery) Where(ps ...predicate.Contagious) *ContagiousQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ContagiousQuery) Limit(limit int) *ContagiousQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ContagiousQuery) Offset(offset int) *ContagiousQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *ContagiousQuery) Order(o ...OrderFunc) *ContagiousQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryStatistic chains the current query on the statistic edge.
func (cq *ContagiousQuery) QueryStatistic() *StatisticQuery {
	query := &StatisticQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contagious.Table, contagious.FieldID, cq.sqlQuery()),
			sqlgraph.To(statistic.Table, statistic.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contagious.StatisticTable, contagious.StatisticColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Contagious entity in the query. Returns *NotFoundError when no contagious was found.
func (cq *ContagiousQuery) First(ctx context.Context) (*Contagious, error) {
	cs, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cs) == 0 {
		return nil, &NotFoundError{contagious.Label}
	}
	return cs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ContagiousQuery) FirstX(ctx context.Context) *Contagious {
	c, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return c
}

// FirstID returns the first Contagious id in the query. Returns *NotFoundError when no id was found.
func (cq *ContagiousQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contagious.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cq *ContagiousQuery) FirstXID(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Contagious entity in the query, returns an error if not exactly one entity was returned.
func (cq *ContagiousQuery) Only(ctx context.Context) (*Contagious, error) {
	cs, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cs) {
	case 1:
		return cs[0], nil
	case 0:
		return nil, &NotFoundError{contagious.Label}
	default:
		return nil, &NotSingularError{contagious.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ContagiousQuery) OnlyX(ctx context.Context) *Contagious {
	c, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// OnlyID returns the only Contagious id in the query, returns an error if not exactly one id was returned.
func (cq *ContagiousQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = &NotSingularError{contagious.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ContagiousQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContagiousSlice.
func (cq *ContagiousQuery) All(ctx context.Context) ([]*Contagious, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ContagiousQuery) AllX(ctx context.Context) []*Contagious {
	cs, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// IDs executes the query and returns a list of Contagious ids.
func (cq *ContagiousQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(contagious.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ContagiousQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ContagiousQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ContagiousQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ContagiousQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ContagiousQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ContagiousQuery) Clone() *ContagiousQuery {
	return &ContagiousQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		unique:     append([]string{}, cq.unique...),
		predicates: append([]predicate.Contagious{}, cq.predicates...),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

//  WithStatistic tells the query-builder to eager-loads the nodes that are connected to
// the "statistic" edge. The optional arguments used to configure the query builder of the edge.
func (cq *ContagiousQuery) WithStatistic(opts ...func(*StatisticQuery)) *ContagiousQuery {
	query := &StatisticQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withStatistic = query
	return cq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Contagious.Query().
//		GroupBy(contagious.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *ContagiousQuery) GroupBy(field string, fields ...string) *ContagiousGroupBy {
	group := &ContagiousGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//	}
//
//	client.Contagious.Query().
//		Select(contagious.FieldName).
//		Scan(ctx, &v)
//
func (cq *ContagiousQuery) Select(field string, fields ...string) *ContagiousSelect {
	selector := &ContagiousSelect{config: cq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(), nil
	}
	return selector
}

func (cq *ContagiousQuery) prepareQuery(ctx context.Context) error {
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ContagiousQuery) sqlAll(ctx context.Context) ([]*Contagious, error) {
	var (
		nodes       = []*Contagious{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withStatistic != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Contagious{config: cq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cq.withStatistic; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Contagious)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Statistic(func(s *sql.Selector) {
			s.Where(sql.InValues(contagious.StatisticColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.contagious_statistic
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "contagious_statistic" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "contagious_statistic" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Statistic = append(node.Edges.Statistic, n)
		}
	}

	return nodes, nil
}

func (cq *ContagiousQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ContagiousQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *ContagiousQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contagious.Table,
			Columns: contagious.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contagious.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ContagiousQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(contagious.Table)
	selector := builder.Select(t1.Columns(contagious.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(contagious.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ContagiousGroupBy is the builder for group-by Contagious entities.
type ContagiousGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ContagiousGroupBy) Aggregate(fns ...AggregateFunc) *ContagiousGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cgb *ContagiousGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *ContagiousGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ContagiousGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *ContagiousGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cgb *ContagiousGroupBy) StringX(ctx context.Context) string {
	v, err := cgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ContagiousGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *ContagiousGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cgb *ContagiousGroupBy) IntX(ctx context.Context) int {
	v, err := cgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ContagiousGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *ContagiousGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cgb *ContagiousGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ContagiousGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *ContagiousGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (cgb *ContagiousGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cgb *ContagiousGroupBy) BoolX(ctx context.Context) bool {
	v, err := cgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *ContagiousGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cgb.sqlQuery().Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ContagiousGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// ContagiousSelect is the builder for select fields of Contagious entities.
type ContagiousSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (cs *ContagiousSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := cs.path(ctx)
	if err != nil {
		return err
	}
	cs.sql = query
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *ContagiousSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ContagiousSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *ContagiousSelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cs *ContagiousSelect) StringX(ctx context.Context) string {
	v, err := cs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ContagiousSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *ContagiousSelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cs *ContagiousSelect) IntX(ctx context.Context) int {
	v, err := cs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ContagiousSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *ContagiousSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cs *ContagiousSelect) Float64X(ctx context.Context) float64 {
	v, err := cs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ContagiousSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *ContagiousSelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (cs *ContagiousSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{contagious.Label}
	default:
		err = fmt.Errorf("ent: ContagiousSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cs *ContagiousSelect) BoolX(ctx context.Context) bool {
	v, err := cs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *ContagiousSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *ContagiousSelect) sqlQuery() sql.Querier {
	selector := cs.sql
	selector.Select(selector.Columns(cs.fields...)...)
	return selector
}
