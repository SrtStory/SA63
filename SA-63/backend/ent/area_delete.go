// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/SrtStory/app/ent/area"
	"github.com/SrtStory/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AreaDelete is the builder for deleting a Area entity.
type AreaDelete struct {
	config
	hooks      []Hook
	mutation   *AreaMutation
	predicates []predicate.Area
}

// Where adds a new predicate to the delete builder.
func (ad *AreaDelete) Where(ps ...predicate.Area) *AreaDelete {
	ad.predicates = append(ad.predicates, ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AreaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AreaDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AreaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: area.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: area.FieldID,
			},
		},
	}
	if ps := ad.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// AreaDeleteOne is the builder for deleting a single Area entity.
type AreaDeleteOne struct {
	ad *AreaDelete
}

// Exec executes the deletion query.
func (ado *AreaDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{area.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AreaDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
