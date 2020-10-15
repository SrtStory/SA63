package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Carrier schema.
type Carrier struct {
	ent.Schema
}

// Fields of the Carrier.
func (Carrier) Fields() []ent.Field {
	return []ent.Field{
			field.String("Name").NotEmpty(),
	}
}

// Edges of the User.
func (Carrier) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("statistic", Statistic.Type),
		
	}	
}
