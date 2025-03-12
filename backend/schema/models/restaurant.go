package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Restaurant struct {
	ent.Schema
}

func (Restaurant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("address"),
	}
}

func (Restaurant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("restaurants").
			Required().
			Unique(),
	}
}
