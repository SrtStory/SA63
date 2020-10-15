package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

//Area Schema.
type Area struct {
	ent.Schema
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
	}
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("statistic", Statistic.Type),
		
	}	
}
