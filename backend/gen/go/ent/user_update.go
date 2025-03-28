// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/credential"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/predicate"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/refreshtoken"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/restaurant"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// ClearName clears the value of the "name" field.
func (uu *UserUpdate) ClearName() *UserUpdate {
	uu.mutation.ClearName()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetFinishedRegistration sets the "finished_registration" field.
func (uu *UserUpdate) SetFinishedRegistration(b bool) *UserUpdate {
	uu.mutation.SetFinishedRegistration(b)
	return uu
}

// SetNillableFinishedRegistration sets the "finished_registration" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFinishedRegistration(b *bool) *UserUpdate {
	if b != nil {
		uu.SetFinishedRegistration(*b)
	}
	return uu
}

// SetLastSignInAt sets the "last_sign_in_at" field.
func (uu *UserUpdate) SetLastSignInAt(t time.Time) *UserUpdate {
	uu.mutation.SetLastSignInAt(t)
	return uu
}

// SetNillableLastSignInAt sets the "last_sign_in_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastSignInAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastSignInAt(*t)
	}
	return uu
}

// ClearLastSignInAt clears the value of the "last_sign_in_at" field.
func (uu *UserUpdate) ClearLastSignInAt() *UserUpdate {
	uu.mutation.ClearLastSignInAt()
	return uu
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by IDs.
func (uu *UserUpdate) AddCredentialIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCredentialIDs(ids...)
	return uu
}

// AddCredentials adds the "credentials" edges to the Credential entity.
func (uu *UserUpdate) AddCredentials(c ...*Credential) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCredentialIDs(ids...)
}

// AddAccessTokenIDs adds the "access_tokens" edge to the RefreshToken entity by IDs.
func (uu *UserUpdate) AddAccessTokenIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAccessTokenIDs(ids...)
	return uu
}

// AddAccessTokens adds the "access_tokens" edges to the RefreshToken entity.
func (uu *UserUpdate) AddAccessTokens(r ...*RefreshToken) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddAccessTokenIDs(ids...)
}

// AddRestaurantIDs adds the "restaurants" edge to the Restaurant entity by IDs.
func (uu *UserUpdate) AddRestaurantIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRestaurantIDs(ids...)
	return uu
}

// AddRestaurants adds the "restaurants" edges to the Restaurant entity.
func (uu *UserUpdate) AddRestaurants(r ...*Restaurant) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRestaurantIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCredentials clears all "credentials" edges to the Credential entity.
func (uu *UserUpdate) ClearCredentials() *UserUpdate {
	uu.mutation.ClearCredentials()
	return uu
}

// RemoveCredentialIDs removes the "credentials" edge to Credential entities by IDs.
func (uu *UserUpdate) RemoveCredentialIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCredentialIDs(ids...)
	return uu
}

// RemoveCredentials removes "credentials" edges to Credential entities.
func (uu *UserUpdate) RemoveCredentials(c ...*Credential) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCredentialIDs(ids...)
}

// ClearAccessTokens clears all "access_tokens" edges to the RefreshToken entity.
func (uu *UserUpdate) ClearAccessTokens() *UserUpdate {
	uu.mutation.ClearAccessTokens()
	return uu
}

// RemoveAccessTokenIDs removes the "access_tokens" edge to RefreshToken entities by IDs.
func (uu *UserUpdate) RemoveAccessTokenIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAccessTokenIDs(ids...)
	return uu
}

// RemoveAccessTokens removes "access_tokens" edges to RefreshToken entities.
func (uu *UserUpdate) RemoveAccessTokens(r ...*RefreshToken) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveAccessTokenIDs(ids...)
}

// ClearRestaurants clears all "restaurants" edges to the Restaurant entity.
func (uu *UserUpdate) ClearRestaurants() *UserUpdate {
	uu.mutation.ClearRestaurants()
	return uu
}

// RemoveRestaurantIDs removes the "restaurants" edge to Restaurant entities by IDs.
func (uu *UserUpdate) RemoveRestaurantIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRestaurantIDs(ids...)
	return uu
}

// RemoveRestaurants removes "restaurants" edges to Restaurant entities.
func (uu *UserUpdate) RemoveRestaurants(r ...*Restaurant) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRestaurantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if uu.mutation.NameCleared() {
		_spec.ClearField(user.FieldName, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uu.mutation.FinishedRegistration(); ok {
		_spec.SetField(user.FieldFinishedRegistration, field.TypeBool, value)
	}
	if value, ok := uu.mutation.LastSignInAt(); ok {
		_spec.SetField(user.FieldLastSignInAt, field.TypeTime, value)
	}
	if uu.mutation.LastSignInAtCleared() {
		_spec.ClearField(user.FieldLastSignInAt, field.TypeTime)
	}
	if uu.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCredentialsIDs(); len(nodes) > 0 && !uu.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: []string{user.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAccessTokensIDs(); len(nodes) > 0 && !uu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: []string{user.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: []string{user.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RestaurantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestaurantsTable,
			Columns: []string{user.RestaurantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(restaurant.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRestaurantsIDs(); len(nodes) > 0 && !uu.mutation.RestaurantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestaurantsTable,
			Columns: []string{user.RestaurantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(restaurant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RestaurantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestaurantsTable,
			Columns: []string{user.RestaurantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(restaurant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// ClearName clears the value of the "name" field.
func (uuo *UserUpdateOne) ClearName() *UserUpdateOne {
	uuo.mutation.ClearName()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetFinishedRegistration sets the "finished_registration" field.
func (uuo *UserUpdateOne) SetFinishedRegistration(b bool) *UserUpdateOne {
	uuo.mutation.SetFinishedRegistration(b)
	return uuo
}

// SetNillableFinishedRegistration sets the "finished_registration" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFinishedRegistration(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetFinishedRegistration(*b)
	}
	return uuo
}

// SetLastSignInAt sets the "last_sign_in_at" field.
func (uuo *UserUpdateOne) SetLastSignInAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastSignInAt(t)
	return uuo
}

// SetNillableLastSignInAt sets the "last_sign_in_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastSignInAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastSignInAt(*t)
	}
	return uuo
}

// ClearLastSignInAt clears the value of the "last_sign_in_at" field.
func (uuo *UserUpdateOne) ClearLastSignInAt() *UserUpdateOne {
	uuo.mutation.ClearLastSignInAt()
	return uuo
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by IDs.
func (uuo *UserUpdateOne) AddCredentialIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCredentialIDs(ids...)
	return uuo
}

// AddCredentials adds the "credentials" edges to the Credential entity.
func (uuo *UserUpdateOne) AddCredentials(c ...*Credential) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCredentialIDs(ids...)
}

// AddAccessTokenIDs adds the "access_tokens" edge to the RefreshToken entity by IDs.
func (uuo *UserUpdateOne) AddAccessTokenIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAccessTokenIDs(ids...)
	return uuo
}

// AddAccessTokens adds the "access_tokens" edges to the RefreshToken entity.
func (uuo *UserUpdateOne) AddAccessTokens(r ...*RefreshToken) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddAccessTokenIDs(ids...)
}

// AddRestaurantIDs adds the "restaurants" edge to the Restaurant entity by IDs.
func (uuo *UserUpdateOne) AddRestaurantIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRestaurantIDs(ids...)
	return uuo
}

// AddRestaurants adds the "restaurants" edges to the Restaurant entity.
func (uuo *UserUpdateOne) AddRestaurants(r ...*Restaurant) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRestaurantIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCredentials clears all "credentials" edges to the Credential entity.
func (uuo *UserUpdateOne) ClearCredentials() *UserUpdateOne {
	uuo.mutation.ClearCredentials()
	return uuo
}

// RemoveCredentialIDs removes the "credentials" edge to Credential entities by IDs.
func (uuo *UserUpdateOne) RemoveCredentialIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCredentialIDs(ids...)
	return uuo
}

// RemoveCredentials removes "credentials" edges to Credential entities.
func (uuo *UserUpdateOne) RemoveCredentials(c ...*Credential) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCredentialIDs(ids...)
}

// ClearAccessTokens clears all "access_tokens" edges to the RefreshToken entity.
func (uuo *UserUpdateOne) ClearAccessTokens() *UserUpdateOne {
	uuo.mutation.ClearAccessTokens()
	return uuo
}

// RemoveAccessTokenIDs removes the "access_tokens" edge to RefreshToken entities by IDs.
func (uuo *UserUpdateOne) RemoveAccessTokenIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAccessTokenIDs(ids...)
	return uuo
}

// RemoveAccessTokens removes "access_tokens" edges to RefreshToken entities.
func (uuo *UserUpdateOne) RemoveAccessTokens(r ...*RefreshToken) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveAccessTokenIDs(ids...)
}

// ClearRestaurants clears all "restaurants" edges to the Restaurant entity.
func (uuo *UserUpdateOne) ClearRestaurants() *UserUpdateOne {
	uuo.mutation.ClearRestaurants()
	return uuo
}

// RemoveRestaurantIDs removes the "restaurants" edge to Restaurant entities by IDs.
func (uuo *UserUpdateOne) RemoveRestaurantIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRestaurantIDs(ids...)
	return uuo
}

// RemoveRestaurants removes "restaurants" edges to Restaurant entities.
func (uuo *UserUpdateOne) RemoveRestaurants(r ...*Restaurant) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRestaurantIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if uuo.mutation.NameCleared() {
		_spec.ClearField(user.FieldName, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uuo.mutation.FinishedRegistration(); ok {
		_spec.SetField(user.FieldFinishedRegistration, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.LastSignInAt(); ok {
		_spec.SetField(user.FieldLastSignInAt, field.TypeTime, value)
	}
	if uuo.mutation.LastSignInAtCleared() {
		_spec.ClearField(user.FieldLastSignInAt, field.TypeTime)
	}
	if uuo.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCredentialsIDs(); len(nodes) > 0 && !uuo.mutation.CredentialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: []string{user.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAccessTokensIDs(); len(nodes) > 0 && !uuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: []string{user.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: []string{user.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RestaurantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestaurantsTable,
			Columns: []string{user.RestaurantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(restaurant.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRestaurantsIDs(); len(nodes) > 0 && !uuo.mutation.RestaurantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestaurantsTable,
			Columns: []string{user.RestaurantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(restaurant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RestaurantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestaurantsTable,
			Columns: []string{user.RestaurantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(restaurant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
