package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Contagious schema.
type Contagious struct {
	ent.Schema
}

// Fields of the Contagious.
func (Contagious) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
	}
}

// Edges of the Contagious.
func (Contagious) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("statistic", Statistic.Type),
	
	}	
}
