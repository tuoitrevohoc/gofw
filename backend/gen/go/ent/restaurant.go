// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/restaurant"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
)

// Restaurant is the model entity for the Restaurant schema.
type Restaurant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RestaurantQuery when eager-loading is set.
	Edges            RestaurantEdges `json:"edges"`
	user_restaurants *int
	selectValues     sql.SelectValues
}

// RestaurantEdges holds the relations/edges for other nodes in the graph.
type RestaurantEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RestaurantEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Restaurant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case restaurant.FieldID:
			values[i] = new(sql.NullInt64)
		case restaurant.FieldName, restaurant.FieldAddress:
			values[i] = new(sql.NullString)
		case restaurant.ForeignKeys[0]: // user_restaurants
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Restaurant fields.
func (r *Restaurant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case restaurant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case restaurant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case restaurant.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				r.Address = value.String
			}
		case restaurant.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_restaurants", value)
			} else if value.Valid {
				r.user_restaurants = new(int)
				*r.user_restaurants = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Restaurant.
// This includes values selected through modifiers, order, etc.
func (r *Restaurant) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Restaurant entity.
func (r *Restaurant) QueryOwner() *UserQuery {
	return NewRestaurantClient(r.config).QueryOwner(r)
}

// Update returns a builder for updating this Restaurant.
// Note that you need to call Restaurant.Unwrap() before calling this method if this Restaurant
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Restaurant) Update() *RestaurantUpdateOne {
	return NewRestaurantClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Restaurant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Restaurant) Unwrap() *Restaurant {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Restaurant is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Restaurant) String() string {
	var builder strings.Builder
	builder.WriteString("Restaurant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(r.Address)
	builder.WriteByte(')')
	return builder.String()
}

// Restaurants is a parsable slice of Restaurant.
type Restaurants []*Restaurant
