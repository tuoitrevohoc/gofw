package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Credential struct {
	ent.Schema
}

func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.String("public_key").NotEmpty(),
		field.String("data").NotEmpty(),
	}
}

func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("credentials").Unique().Required(),
	}
}
