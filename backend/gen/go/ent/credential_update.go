// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/credential"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/predicate"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
)

// CredentialUpdate is the builder for updating Credential entities.
type CredentialUpdate struct {
	config
	hooks    []Hook
	mutation *CredentialMutation
}

// Where appends a list predicates to the CredentialUpdate builder.
func (cu *CredentialUpdate) Where(ps ...predicate.Credential) *CredentialUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetPublicKey sets the "public_key" field.
func (cu *CredentialUpdate) SetPublicKey(s string) *CredentialUpdate {
	cu.mutation.SetPublicKey(s)
	return cu
}

// SetNillablePublicKey sets the "public_key" field if the given value is not nil.
func (cu *CredentialUpdate) SetNillablePublicKey(s *string) *CredentialUpdate {
	if s != nil {
		cu.SetPublicKey(*s)
	}
	return cu
}

// SetData sets the "data" field.
func (cu *CredentialUpdate) SetData(s string) *CredentialUpdate {
	cu.mutation.SetData(s)
	return cu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (cu *CredentialUpdate) SetNillableData(s *string) *CredentialUpdate {
	if s != nil {
		cu.SetData(*s)
	}
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CredentialUpdate) SetUserID(id int) *CredentialUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CredentialUpdate) SetUser(u *User) *CredentialUpdate {
	return cu.SetUserID(u.ID)
}

// Mutation returns the CredentialMutation object of the builder.
func (cu *CredentialUpdate) Mutation() *CredentialMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CredentialUpdate) ClearUser() *CredentialUpdate {
	cu.mutation.ClearUser()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CredentialUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CredentialUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CredentialUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CredentialUpdate) check() error {
	if v, ok := cu.mutation.PublicKey(); ok {
		if err := credential.PublicKeyValidator(v); err != nil {
			return &ValidationError{Name: "public_key", err: fmt.Errorf(`ent: validator failed for field "Credential.public_key": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Data(); ok {
		if err := credential.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "Credential.data": %w`, err)}
		}
	}
	if cu.mutation.UserCleared() && len(cu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Credential.user"`)
	}
	return nil
}

func (cu *CredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(credential.Table, credential.Columns, sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.PublicKey(); ok {
		_spec.SetField(credential.FieldPublicKey, field.TypeString, value)
	}
	if value, ok := cu.mutation.Data(); ok {
		_spec.SetField(credential.FieldData, field.TypeString, value)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.UserTable,
			Columns: []string{credential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.UserTable,
			Columns: []string{credential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CredentialUpdateOne is the builder for updating a single Credential entity.
type CredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CredentialMutation
}

// SetPublicKey sets the "public_key" field.
func (cuo *CredentialUpdateOne) SetPublicKey(s string) *CredentialUpdateOne {
	cuo.mutation.SetPublicKey(s)
	return cuo
}

// SetNillablePublicKey sets the "public_key" field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillablePublicKey(s *string) *CredentialUpdateOne {
	if s != nil {
		cuo.SetPublicKey(*s)
	}
	return cuo
}

// SetData sets the "data" field.
func (cuo *CredentialUpdateOne) SetData(s string) *CredentialUpdateOne {
	cuo.mutation.SetData(s)
	return cuo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableData(s *string) *CredentialUpdateOne {
	if s != nil {
		cuo.SetData(*s)
	}
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CredentialUpdateOne) SetUserID(id int) *CredentialUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CredentialUpdateOne) SetUser(u *User) *CredentialUpdateOne {
	return cuo.SetUserID(u.ID)
}

// Mutation returns the CredentialMutation object of the builder.
func (cuo *CredentialUpdateOne) Mutation() *CredentialMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CredentialUpdateOne) ClearUser() *CredentialUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// Where appends a list predicates to the CredentialUpdate builder.
func (cuo *CredentialUpdateOne) Where(ps ...predicate.Credential) *CredentialUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CredentialUpdateOne) Select(field string, fields ...string) *CredentialUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Credential entity.
func (cuo *CredentialUpdateOne) Save(ctx context.Context) (*Credential, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CredentialUpdateOne) SaveX(ctx context.Context) *Credential {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CredentialUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CredentialUpdateOne) check() error {
	if v, ok := cuo.mutation.PublicKey(); ok {
		if err := credential.PublicKeyValidator(v); err != nil {
			return &ValidationError{Name: "public_key", err: fmt.Errorf(`ent: validator failed for field "Credential.public_key": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Data(); ok {
		if err := credential.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "Credential.data": %w`, err)}
		}
	}
	if cuo.mutation.UserCleared() && len(cuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Credential.user"`)
	}
	return nil
}

func (cuo *CredentialUpdateOne) sqlSave(ctx context.Context) (_node *Credential, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(credential.Table, credential.Columns, sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Credential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, credential.FieldID)
		for _, f := range fields {
			if !credential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != credential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.PublicKey(); ok {
		_spec.SetField(credential.FieldPublicKey, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Data(); ok {
		_spec.SetField(credential.FieldData, field.TypeString, value)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.UserTable,
			Columns: []string{credential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.UserTable,
			Columns: []string{credential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Credential{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
