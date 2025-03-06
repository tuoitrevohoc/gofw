package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("avatar").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
