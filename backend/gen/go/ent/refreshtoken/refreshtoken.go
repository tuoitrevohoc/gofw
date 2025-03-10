// Code generated by ent, DO NOT EDIT.

package refreshtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the refreshtoken type in the database.
	Label = "refresh_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldRefreshAt holds the string denoting the refresh_at field in the database.
	FieldRefreshAt = "refresh_at"
	// FieldExpireAt holds the string denoting the expire_at field in the database.
	FieldExpireAt = "expire_at"
	// FieldIPAddress holds the string denoting the ip_address field in the database.
	FieldIPAddress = "ip_address"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldUserAgent holds the string denoting the user_agent field in the database.
	FieldUserAgent = "user_agent"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the refreshtoken in the database.
	Table = "refresh_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "refresh_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_access_tokens"
)

// Columns holds all SQL columns for refreshtoken fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldCreatedAt,
	FieldRefreshAt,
	FieldExpireAt,
	FieldIPAddress,
	FieldIsActive,
	FieldUserAgent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "refresh_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_access_tokens",
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
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultRefreshAt holds the default value on creation for the "refresh_at" field.
	DefaultRefreshAt func() time.Time
	// IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	IPAddressValidator func(string) error
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// UserAgentValidator is a validator for the "user_agent" field. It is called by the builders before save.
	UserAgentValidator func(string) error
)

// OrderOption defines the ordering options for the RefreshToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByRefreshAt orders the results by the refresh_at field.
func ByRefreshAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshAt, opts...).ToFunc()
}

// ByExpireAt orders the results by the expire_at field.
func ByExpireAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpireAt, opts...).ToFunc()
}

// ByIPAddress orders the results by the ip_address field.
func ByIPAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPAddress, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByUserAgent orders the results by the user_agent field.
func ByUserAgent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserAgent, opts...).ToFunc()
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
