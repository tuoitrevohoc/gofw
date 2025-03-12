package models

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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

func (Restaurant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("restaurants"),
		entgql.RelayConnection(),
	}
}
