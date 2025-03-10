// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/credential"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/refreshtoken"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CredentialQuery) CollectFields(ctx context.Context, satisfies ...string) (*CredentialQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CredentialQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(credential.Columns))
		selectedFields = []string{credential.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			c.withUser = query
		case "publicKey":
			if _, ok := fieldSeen[credential.FieldPublicKey]; !ok {
				selectedFields = append(selectedFields, credential.FieldPublicKey)
				fieldSeen[credential.FieldPublicKey] = struct{}{}
			}
		case "data":
			if _, ok := fieldSeen[credential.FieldData]; !ok {
				selectedFields = append(selectedFields, credential.FieldData)
				fieldSeen[credential.FieldData] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type credentialPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CredentialPaginateOption
}

func newCredentialPaginateArgs(rv map[string]any) *credentialPaginateArgs {
	args := &credentialPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (rt *RefreshTokenQuery) CollectFields(ctx context.Context, satisfies ...string) (*RefreshTokenQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return rt, nil
	}
	if err := rt.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return rt, nil
}

func (rt *RefreshTokenQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(refreshtoken.Columns))
		selectedFields = []string{refreshtoken.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: rt.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			rt.withUser = query
		case "token":
			if _, ok := fieldSeen[refreshtoken.FieldToken]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldToken)
				fieldSeen[refreshtoken.FieldToken] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[refreshtoken.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldCreatedAt)
				fieldSeen[refreshtoken.FieldCreatedAt] = struct{}{}
			}
		case "refreshAt":
			if _, ok := fieldSeen[refreshtoken.FieldRefreshAt]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldRefreshAt)
				fieldSeen[refreshtoken.FieldRefreshAt] = struct{}{}
			}
		case "expireAt":
			if _, ok := fieldSeen[refreshtoken.FieldExpireAt]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldExpireAt)
				fieldSeen[refreshtoken.FieldExpireAt] = struct{}{}
			}
		case "ipAddress":
			if _, ok := fieldSeen[refreshtoken.FieldIPAddress]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldIPAddress)
				fieldSeen[refreshtoken.FieldIPAddress] = struct{}{}
			}
		case "isActive":
			if _, ok := fieldSeen[refreshtoken.FieldIsActive]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldIsActive)
				fieldSeen[refreshtoken.FieldIsActive] = struct{}{}
			}
		case "userAgent":
			if _, ok := fieldSeen[refreshtoken.FieldUserAgent]; !ok {
				selectedFields = append(selectedFields, refreshtoken.FieldUserAgent)
				fieldSeen[refreshtoken.FieldUserAgent] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		rt.Select(selectedFields...)
	}
	return nil
}

type refreshtokenPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RefreshTokenPaginateOption
}

func newRefreshTokenPaginateArgs(rv map[string]any) *refreshtokenPaginateArgs {
	args := &refreshtokenPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "credentials":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CredentialClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, credentialImplementors)...); err != nil {
				return err
			}
			u.WithNamedCredentials(alias, func(wq *CredentialQuery) {
				*wq = *query
			})

		case "accessTokens":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RefreshTokenClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, refreshtokenImplementors)...); err != nil {
				return err
			}
			u.WithNamedAccessTokens(alias, func(wq *RefreshTokenQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "password":
			if _, ok := fieldSeen[user.FieldPassword]; !ok {
				selectedFields = append(selectedFields, user.FieldPassword)
				fieldSeen[user.FieldPassword] = struct{}{}
			}
		case "avatar":
			if _, ok := fieldSeen[user.FieldAvatar]; !ok {
				selectedFields = append(selectedFields, user.FieldAvatar)
				fieldSeen[user.FieldAvatar] = struct{}{}
			}
		case "finishedRegistration":
			if _, ok := fieldSeen[user.FieldFinishedRegistration]; !ok {
				selectedFields = append(selectedFields, user.FieldFinishedRegistration)
				fieldSeen[user.FieldFinishedRegistration] = struct{}{}
			}
		case "lastSignInAt":
			if _, ok := fieldSeen[user.FieldLastSignInAt]; !ok {
				selectedFields = append(selectedFields, user.FieldLastSignInAt)
				fieldSeen[user.FieldLastSignInAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
