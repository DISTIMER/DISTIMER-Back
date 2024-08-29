// Code generated by ent, DO NOT EDIT.

package fcmtoken

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the fcmtoken type in the database.
	Label = "fcm_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPushToken holds the string denoting the push_token field in the database.
	FieldPushToken = "push_token"
	// EdgeSession holds the string denoting the session edge name in mutations.
	EdgeSession = "session"
	// Table holds the table name of the fcmtoken in the database.
	Table = "fcm_tokens"
	// SessionTable is the table that holds the session relation/edge.
	SessionTable = "fcm_tokens"
	// SessionInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionInverseTable = "sessions"
	// SessionColumn is the table column denoting the session relation/edge.
	SessionColumn = "session_fcm_token"
)

// Columns holds all SQL columns for fcmtoken fields.
var Columns = []string{
	FieldID,
	FieldPushToken,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "fcm_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"session_fcm_token",
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

// OrderOption defines the ordering options for the FCMToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPushToken orders the results by the push_token field.
func ByPushToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPushToken, opts...).ToFunc()
}

// BySessionField orders the results by session field.
func BySessionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSessionStep(), sql.OrderByField(field, opts...))
	}
}
func newSessionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SessionTable, SessionColumn),
	)
}
