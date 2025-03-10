package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// store the refresh token for the user

type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("refresh_at").Default(time.Now),
		field.Time("expire_at"),
		field.String("ip_address").NotEmpty(),
		field.Bool("is_active").Default(true),
		field.String("user_agent").NotEmpty(),
	}
}

func (RefreshToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token"),
	}
}

func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("access_tokens").Unique().Required(),
	}
}
