// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/SrtStory/app/ent/area"
	"github.com/SrtStory/app/ent/carrier"
	"github.com/SrtStory/app/ent/contagious"
	"github.com/SrtStory/app/ent/employee"
	"github.com/SrtStory/app/ent/patient"
	"github.com/SrtStory/app/ent/predicate"
	"github.com/SrtStory/app/ent/statistic"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// StatisticUpdate is the builder for updating Statistic entities.
type StatisticUpdate struct {
	config
	hooks      []Hook
	mutation   *StatisticMutation
	predicates []predicate.Statistic
}

// Where adds a new predicate for the builder.
func (su *StatisticUpdate) Where(ps ...predicate.Statistic) *StatisticUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetEmployeeID sets the employee edge to Employee by id.
func (su *StatisticUpdate) SetEmployeeID(id int) *StatisticUpdate {
	su.mutation.SetEmployeeID(id)
	return su
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (su *StatisticUpdate) SetNillableEmployeeID(id *int) *StatisticUpdate {
	if id != nil {
		su = su.SetEmployeeID(*id)
	}
	return su
}

// SetEmployee sets the employee edge to Employee.
func (su *StatisticUpdate) SetEmployee(e *Employee) *StatisticUpdate {
	return su.SetEmployeeID(e.ID)
}

// SetContagiousID sets the contagious edge to Contagious by id.
func (su *StatisticUpdate) SetContagiousID(id int) *StatisticUpdate {
	su.mutation.SetContagiousID(id)
	return su
}

// SetNillableContagiousID sets the contagious edge to Contagious by id if the given value is not nil.
func (su *StatisticUpdate) SetNillableContagiousID(id *int) *StatisticUpdate {
	if id != nil {
		su = su.SetContagiousID(*id)
	}
	return su
}

// SetContagious sets the contagious edge to Contagious.
func (su *StatisticUpdate) SetContagious(c *Contagious) *StatisticUpdate {
	return su.SetContagiousID(c.ID)
}

// SetCarrierID sets the carrier edge to Carrier by id.
func (su *StatisticUpdate) SetCarrierID(id int) *StatisticUpdate {
	su.mutation.SetCarrierID(id)
	return su
}

// SetNillableCarrierID sets the carrier edge to Carrier by id if the given value is not nil.
func (su *StatisticUpdate) SetNillableCarrierID(id *int) *StatisticUpdate {
	if id != nil {
		su = su.SetCarrierID(*id)
	}
	return su
}

// SetCarrier sets the carrier edge to Carrier.
func (su *StatisticUpdate) SetCarrier(c *Carrier) *StatisticUpdate {
	return su.SetCarrierID(c.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (su *StatisticUpdate) SetPatientID(id int) *StatisticUpdate {
	su.mutation.SetPatientID(id)
	return su
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (su *StatisticUpdate) SetNillablePatientID(id *int) *StatisticUpdate {
	if id != nil {
		su = su.SetPatientID(*id)
	}
	return su
}

// SetPatient sets the patient edge to Patient.
func (su *StatisticUpdate) SetPatient(p *Patient) *StatisticUpdate {
	return su.SetPatientID(p.ID)
}

// SetAreaID sets the area edge to Area by id.
func (su *StatisticUpdate) SetAreaID(id int) *StatisticUpdate {
	su.mutation.SetAreaID(id)
	return su
}

// SetNillableAreaID sets the area edge to Area by id if the given value is not nil.
func (su *StatisticUpdate) SetNillableAreaID(id *int) *StatisticUpdate {
	if id != nil {
		su = su.SetAreaID(*id)
	}
	return su
}

// SetArea sets the area edge to Area.
func (su *StatisticUpdate) SetArea(a *Area) *StatisticUpdate {
	return su.SetAreaID(a.ID)
}

// Mutation returns the StatisticMutation object of the builder.
func (su *StatisticUpdate) Mutation() *StatisticMutation {
	return su.mutation
}

// ClearEmployee clears the employee edge to Employee.
func (su *StatisticUpdate) ClearEmployee() *StatisticUpdate {
	su.mutation.ClearEmployee()
	return su
}

// ClearContagious clears the contagious edge to Contagious.
func (su *StatisticUpdate) ClearContagious() *StatisticUpdate {
	su.mutation.ClearContagious()
	return su
}

// ClearCarrier clears the carrier edge to Carrier.
func (su *StatisticUpdate) ClearCarrier() *StatisticUpdate {
	su.mutation.ClearCarrier()
	return su
}

// ClearPatient clears the patient edge to Patient.
func (su *StatisticUpdate) ClearPatient() *StatisticUpdate {
	su.mutation.ClearPatient()
	return su
}

// ClearArea clears the area edge to Area.
func (su *StatisticUpdate) ClearArea() *StatisticUpdate {
	su.mutation.ClearArea()
	return su
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *StatisticUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatisticMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StatisticUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatisticUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatisticUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StatisticUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statistic.Table,
			Columns: statistic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statistic.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if su.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.EmployeeTable,
			Columns: []string{statistic.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.EmployeeTable,
			Columns: []string{statistic.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ContagiousCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.ContagiousTable,
			Columns: []string{statistic.ContagiousColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contagious.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ContagiousIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.ContagiousTable,
			Columns: []string{statistic.ContagiousColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contagious.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CarrierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.CarrierTable,
			Columns: []string{statistic.CarrierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrier.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CarrierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.CarrierTable,
			Columns: []string{statistic.CarrierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PatientTable,
			Columns: []string{statistic.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PatientTable,
			Columns: []string{statistic.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.AreaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.AreaTable,
			Columns: []string{statistic.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.AreaTable,
			Columns: []string{statistic.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statistic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StatisticUpdateOne is the builder for updating a single Statistic entity.
type StatisticUpdateOne struct {
	config
	hooks    []Hook
	mutation *StatisticMutation
}

// SetEmployeeID sets the employee edge to Employee by id.
func (suo *StatisticUpdateOne) SetEmployeeID(id int) *StatisticUpdateOne {
	suo.mutation.SetEmployeeID(id)
	return suo
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (suo *StatisticUpdateOne) SetNillableEmployeeID(id *int) *StatisticUpdateOne {
	if id != nil {
		suo = suo.SetEmployeeID(*id)
	}
	return suo
}

// SetEmployee sets the employee edge to Employee.
func (suo *StatisticUpdateOne) SetEmployee(e *Employee) *StatisticUpdateOne {
	return suo.SetEmployeeID(e.ID)
}

// SetContagiousID sets the contagious edge to Contagious by id.
func (suo *StatisticUpdateOne) SetContagiousID(id int) *StatisticUpdateOne {
	suo.mutation.SetContagiousID(id)
	return suo
}

// SetNillableContagiousID sets the contagious edge to Contagious by id if the given value is not nil.
func (suo *StatisticUpdateOne) SetNillableContagiousID(id *int) *StatisticUpdateOne {
	if id != nil {
		suo = suo.SetContagiousID(*id)
	}
	return suo
}

// SetContagious sets the contagious edge to Contagious.
func (suo *StatisticUpdateOne) SetContagious(c *Contagious) *StatisticUpdateOne {
	return suo.SetContagiousID(c.ID)
}

// SetCarrierID sets the carrier edge to Carrier by id.
func (suo *StatisticUpdateOne) SetCarrierID(id int) *StatisticUpdateOne {
	suo.mutation.SetCarrierID(id)
	return suo
}

// SetNillableCarrierID sets the carrier edge to Carrier by id if the given value is not nil.
func (suo *StatisticUpdateOne) SetNillableCarrierID(id *int) *StatisticUpdateOne {
	if id != nil {
		suo = suo.SetCarrierID(*id)
	}
	return suo
}

// SetCarrier sets the carrier edge to Carrier.
func (suo *StatisticUpdateOne) SetCarrier(c *Carrier) *StatisticUpdateOne {
	return suo.SetCarrierID(c.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (suo *StatisticUpdateOne) SetPatientID(id int) *StatisticUpdateOne {
	suo.mutation.SetPatientID(id)
	return suo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (suo *StatisticUpdateOne) SetNillablePatientID(id *int) *StatisticUpdateOne {
	if id != nil {
		suo = suo.SetPatientID(*id)
	}
	return suo
}

// SetPatient sets the patient edge to Patient.
func (suo *StatisticUpdateOne) SetPatient(p *Patient) *StatisticUpdateOne {
	return suo.SetPatientID(p.ID)
}

// SetAreaID sets the area edge to Area by id.
func (suo *StatisticUpdateOne) SetAreaID(id int) *StatisticUpdateOne {
	suo.mutation.SetAreaID(id)
	return suo
}

// SetNillableAreaID sets the area edge to Area by id if the given value is not nil.
func (suo *StatisticUpdateOne) SetNillableAreaID(id *int) *StatisticUpdateOne {
	if id != nil {
		suo = suo.SetAreaID(*id)
	}
	return suo
}

// SetArea sets the area edge to Area.
func (suo *StatisticUpdateOne) SetArea(a *Area) *StatisticUpdateOne {
	return suo.SetAreaID(a.ID)
}

// Mutation returns the StatisticMutation object of the builder.
func (suo *StatisticUpdateOne) Mutation() *StatisticMutation {
	return suo.mutation
}

// ClearEmployee clears the employee edge to Employee.
func (suo *StatisticUpdateOne) ClearEmployee() *StatisticUpdateOne {
	suo.mutation.ClearEmployee()
	return suo
}

// ClearContagious clears the contagious edge to Contagious.
func (suo *StatisticUpdateOne) ClearContagious() *StatisticUpdateOne {
	suo.mutation.ClearContagious()
	return suo
}

// ClearCarrier clears the carrier edge to Carrier.
func (suo *StatisticUpdateOne) ClearCarrier() *StatisticUpdateOne {
	suo.mutation.ClearCarrier()
	return suo
}

// ClearPatient clears the patient edge to Patient.
func (suo *StatisticUpdateOne) ClearPatient() *StatisticUpdateOne {
	suo.mutation.ClearPatient()
	return suo
}

// ClearArea clears the area edge to Area.
func (suo *StatisticUpdateOne) ClearArea() *StatisticUpdateOne {
	suo.mutation.ClearArea()
	return suo
}

// Save executes the query and returns the updated entity.
func (suo *StatisticUpdateOne) Save(ctx context.Context) (*Statistic, error) {

	var (
		err  error
		node *Statistic
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatisticMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StatisticUpdateOne) SaveX(ctx context.Context) *Statistic {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *StatisticUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatisticUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StatisticUpdateOne) sqlSave(ctx context.Context) (s *Statistic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statistic.Table,
			Columns: statistic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statistic.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Statistic.ID for update")}
	}
	_spec.Node.ID.Value = id
	if suo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.EmployeeTable,
			Columns: []string{statistic.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.EmployeeTable,
			Columns: []string{statistic.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ContagiousCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.ContagiousTable,
			Columns: []string{statistic.ContagiousColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contagious.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ContagiousIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.ContagiousTable,
			Columns: []string{statistic.ContagiousColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contagious.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CarrierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.CarrierTable,
			Columns: []string{statistic.CarrierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrier.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CarrierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.CarrierTable,
			Columns: []string{statistic.CarrierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: carrier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PatientTable,
			Columns: []string{statistic.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PatientTable,
			Columns: []string{statistic.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.AreaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.AreaTable,
			Columns: []string{statistic.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.AreaTable,
			Columns: []string{statistic.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Statistic{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statistic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
