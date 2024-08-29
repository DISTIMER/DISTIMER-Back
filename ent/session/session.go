// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRefreshToken holds the string denoting the refresh_token field in the database.
	FieldRefreshToken = "refresh_token"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeviceType holds the string denoting the device_type field in the database.
	FieldDeviceType = "device_type"
	// FieldLastActive holds the string denoting the last_active field in the database.
	FieldLastActive = "last_active"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeApnsToken holds the string denoting the apns_token edge name in mutations.
	EdgeApnsToken = "apns_token"
	// EdgeFcmToken holds the string denoting the fcm_token edge name in mutations.
	EdgeFcmToken = "fcm_token"
	// Table holds the table name of the session in the database.
	Table = "sessions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "sessions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "apns_tokens"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_sessions"
	// ApnsTokenTable is the table that holds the apns_token relation/edge.
	ApnsTokenTable = "ap_ns_tokens"
	// ApnsTokenInverseTable is the table name for the APNsToken entity.
	// It exists in this package in order to avoid circular dependency with the "apnstoken" package.
	ApnsTokenInverseTable = "ap_ns_tokens"
	// ApnsTokenColumn is the table column denoting the apns_token relation/edge.
	ApnsTokenColumn = "session_apns_token"
	// FcmTokenTable is the table that holds the fcm_token relation/edge.
	FcmTokenTable = "fcm_tokens"
	// FcmTokenInverseTable is the table name for the FCMToken entity.
	// It exists in this package in order to avoid circular dependency with the "fcmtoken" package.
	FcmTokenInverseTable = "fcm_tokens"
	// FcmTokenColumn is the table column denoting the fcm_token relation/edge.
	FcmTokenColumn = "session_fcm_token"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldRefreshToken,
	FieldCreatedAt,
	FieldDeviceType,
	FieldLastActive,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sessions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_sessions",
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
	// DefaultRefreshToken holds the default value on creation for the "refresh_token" field.
	DefaultRefreshToken func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastActive holds the default value on creation for the "last_active" field.
	DefaultLastActive func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRefreshToken orders the results by the refresh_token field.
func ByRefreshToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshToken, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDeviceType orders the results by the device_type field.
func ByDeviceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceType, opts...).ToFunc()
}

// ByLastActive orders the results by the last_active field.
func ByLastActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastActive, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByApnsTokenField orders the results by apns_token field.
func ByApnsTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApnsTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByFcmTokenField orders the results by fcm_token field.
func ByFcmTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFcmTokenStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newApnsTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApnsTokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ApnsTokenTable, ApnsTokenColumn),
	)
}
func newFcmTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FcmTokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FcmTokenTable, FcmTokenColumn),
	)
}
