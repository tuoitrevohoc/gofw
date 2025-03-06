package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AuthSession struct {
	ent.Schema
}

func (AuthSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("data").NotEmpty(), // data in json format
		field.Int("user_id").Optional(),
	}
}

func (AuthSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Unique().
			Ref("auth_sessions").
			Field("user_id").
			Unique(),
	}
}

func (AuthSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
