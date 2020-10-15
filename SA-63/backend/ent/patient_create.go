// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/SrtStory/app/ent/patient"
	"github.com/SrtStory/app/ent/statistic"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetGender sets the Gender field.
func (pc *PatientCreate) SetGender(s string) *PatientCreate {
	pc.mutation.SetGender(s)
	return pc
}

// AddStatisticIDs adds the statistic edge to Statistic by ids.
func (pc *PatientCreate) AddStatisticIDs(ids ...int) *PatientCreate {
	pc.mutation.AddStatisticIDs(ids...)
	return pc
}

// AddStatistic adds the statistic edges to Statistic.
func (pc *PatientCreate) AddStatistic(s ...*Statistic) *PatientCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStatisticIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	if _, ok := pc.mutation.Gender(); !ok {
		return nil, &ValidationError{Name: "Gender", err: errors.New("ent: missing required field \"Gender\"")}
	}
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldGender,
		})
		pa.Gender = value
	}
	if nodes := pc.mutation.StatisticIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.StatisticTable,
			Columns: []string{patient.StatisticColumn},
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
	return pa, _spec
}
