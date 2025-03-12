// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "public_key", Type: field.TypeString},
		{Name: "data", Type: field.TypeString},
		{Name: "user_credentials", Type: field.TypeInt},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:       "credentials",
		Columns:    CredentialsColumns,
		PrimaryKey: []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credentials_users_credentials",
				Columns:    []*schema.Column{CredentialsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "refresh_at", Type: field.TypeTime},
		{Name: "expire_at", Type: field.TypeTime},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "user_access_tokens", Type: field.TypeInt},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refresh_tokens_users_access_tokens",
				Columns:    []*schema.Column{RefreshTokensColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "refreshtoken_token",
				Unique:  false,
				Columns: []*schema.Column{RefreshTokensColumns[1]},
			},
		},
	}
	// RestaurantsColumns holds the columns for the "restaurants" table.
	RestaurantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "user_restaurants", Type: field.TypeInt},
	}
	// RestaurantsTable holds the schema information for the "restaurants" table.
	RestaurantsTable = &schema.Table{
		Name:       "restaurants",
		Columns:    RestaurantsColumns,
		PrimaryKey: []*schema.Column{RestaurantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "restaurants_users_restaurants",
				Columns:    []*schema.Column{RestaurantsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "finished_registration", Type: field.TypeBool, Default: false},
		{Name: "last_sign_in_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CredentialsTable,
		RefreshTokensTable,
		RestaurantsTable,
		UsersTable,
	}
)

func init() {
	CredentialsTable.ForeignKeys[0].RefTable = UsersTable
	CredentialsTable.Annotation = &entsql.Annotation{
		IncrementStart: func(i int) *int { return &i }(4294967296),
	}
	RefreshTokensTable.ForeignKeys[0].RefTable = UsersTable
	RefreshTokensTable.Annotation = &entsql.Annotation{
		IncrementStart: func(i int) *int { return &i }(8589934592),
	}
	RestaurantsTable.ForeignKeys[0].RefTable = UsersTable
	RestaurantsTable.Annotation = &entsql.Annotation{
		IncrementStart: func(i int) *int { return &i }(17179869184),
	}
	UsersTable.Annotation = &entsql.Annotation{
		IncrementStart: func(i int) *int { return &i }(12884901888),
	}
}
