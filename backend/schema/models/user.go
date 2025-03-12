package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// UserMixin implements common fields for the User schema.
type UserMixin struct {
	mixin.Schema
}

func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").Optional(),
		field.String("avatar").Optional(),
		field.Bool("finished_registration").Default(false),
		field.Time("last_sign_in_at").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("credentials", Credential.Type),
		edge.To("access_tokens", RefreshToken.Type),
		edge.To("restaurants", Restaurant.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}
