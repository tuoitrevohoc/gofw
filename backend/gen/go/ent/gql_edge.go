// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Credential) User(ctx context.Context) (*User, error) {
	result, err := c.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryUser().Only(ctx)
	}
	return result, err
}

func (rt *RefreshToken) User(ctx context.Context) (*User, error) {
	result, err := rt.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = rt.QueryUser().Only(ctx)
	}
	return result, err
}

func (r *Restaurant) Owner(ctx context.Context) (*User, error) {
	result, err := r.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryOwner().Only(ctx)
	}
	return result, err
}

func (u *User) Credentials(ctx context.Context) (result []*Credential, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCredentials(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CredentialsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCredentials().All(ctx)
	}
	return result, err
}

func (u *User) AccessTokens(ctx context.Context) (result []*RefreshToken, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedAccessTokens(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.AccessTokensOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryAccessTokens().All(ctx)
	}
	return result, err
}

func (u *User) Restaurants(ctx context.Context) (result []*Restaurant, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedRestaurants(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.RestaurantsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryRestaurants().All(ctx)
	}
	return result, err
}
