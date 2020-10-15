// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/SrtStory/app/ent/carrier"
	"github.com/SrtStory/app/ent/statistic"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CarrierCreate is the builder for creating a Carrier entity.
type CarrierCreate struct {
	config
	mutation *CarrierMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (cc *CarrierCreate) SetName(s string) *CarrierCreate {
	cc.mutation.SetName(s)
	return cc
}

// AddStatisticIDs adds the statistic edge to Statistic by ids.
func (cc *CarrierCreate) AddStatisticIDs(ids ...int) *CarrierCreate {
	cc.mutation.AddStatisticIDs(ids...)
	return cc
}

// AddStatistic adds the statistic edges to Statistic.
func (cc *CarrierCreate) AddStatistic(s ...*Statistic) *CarrierCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStatisticIDs(ids...)
}

// Mutation returns the CarrierMutation object of the builder.
func (cc *CarrierCreate) Mutation() *CarrierMutation {
	return cc.mutation
}

// Save creates the Carrier in the database.
func (cc *CarrierCreate) Save(ctx context.Context) (*Carrier, error) {
	if _, ok := cc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := carrier.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Carrier
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarrierMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarrierCreate) SaveX(ctx context.Context) *Carrier {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CarrierCreate) sqlSave(ctx context.Context) (*Carrier, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CarrierCreate) createSpec() (*Carrier, *sqlgraph.CreateSpec) {
	var (
		c     = &Carrier{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: carrier.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: carrier.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: carrier.FieldName,
		})
		c.Name = value
	}
	if nodes := cc.mutation.StatisticIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   carrier.StatisticTable,
			Columns: []string{carrier.StatisticColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statistic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}