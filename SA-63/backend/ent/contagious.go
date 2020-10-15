// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/SrtStory/app/ent/contagious"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Contagious is the model entity for the Contagious schema.
type Contagious struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContagiousQuery when eager-loading is set.
	Edges ContagiousEdges `json:"edges"`
}

// ContagiousEdges holds the relations/edges for other nodes in the graph.
type ContagiousEdges struct {
	// Statistic holds the value of the statistic edge.
	Statistic []*Statistic
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StatisticOrErr returns the Statistic value or an error if the edge
// was not loaded in eager-loading.
func (e ContagiousEdges) StatisticOrErr() ([]*Statistic, error) {
	if e.loadedTypes[0] {
		return e.Statistic, nil
	}
	return nil, &NotLoadedError{edge: "statistic"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contagious) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contagious fields.
func (c *Contagious) assignValues(values ...interface{}) error {
	if m, n := len(values), len(contagious.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		c.Name = value.String
	}
	return nil
}

// QueryStatistic queries the statistic edge of the Contagious.
func (c *Contagious) QueryStatistic() *StatisticQuery {
	return (&ContagiousClient{config: c.config}).QueryStatistic(c)
}

// Update returns a builder for updating this Contagious.
// Note that, you need to call Contagious.Unwrap() before calling this method, if this Contagious
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contagious) Update() *ContagiousUpdateOne {
	return (&ContagiousClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Contagious) Unwrap() *Contagious {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contagious is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contagious) String() string {
	var builder strings.Builder
	builder.WriteString("Contagious(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", Name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ContagiousSlice is a parsable slice of Contagious.
type ContagiousSlice []*Contagious

func (c ContagiousSlice) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
