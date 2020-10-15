package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent"
//	"github.com/facebookincubator/ent/schema/field"
)

//Statistic Schema.
type Statistic struct {
	ent.Schema
}

// Fields of the Statistic.
func (Statistic) Fields() []ent.Field {
	return []ent.Field{
		
	}
}

// Edges of the Statistic.
func (Statistic) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("employee", Employee.Type).Ref("statistic").Unique(),
		edge.From("contagious", Contagious.Type).Ref("statistic").Unique(),
		edge.From("carrier", Carrier.Type).Ref("statistic").Unique(),
		edge.From("patient", Patient.Type).Ref("statistic").Unique(),
		edge.From("area", Area.Type).Ref("statistic").Unique(),
	}	
}
