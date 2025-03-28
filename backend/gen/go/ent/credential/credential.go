// Code generated by ent, DO NOT EDIT.

package credential

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the credential type in the database.
	Label = "credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the credential in the database.
	Table = "credentials"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "credentials"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_credentials"
)

// Columns holds all SQL columns for credential fields.
var Columns = []string{
	FieldID,
	FieldPublicKey,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "credentials"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_credentials",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// PublicKeyValidator is a validator for the "public_key" field. It is called by the builders before save.
	PublicKeyValidator func(string) error
	// DataValidator is a validator for the "data" field. It is called by the builders before save.
	DataValidator func(string) error
)

// OrderOption defines the ordering options for the Credential queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPublicKey orders the results by the public_key field.
func ByPublicKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicKey, opts...).ToFunc()
}

// ByData orders the results by the data field.
func ByData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldData, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
