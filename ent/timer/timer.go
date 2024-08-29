// Code generated by ent, DO NOT EDIT.

package timer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the timer type in the database.
	Label = "timer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldSubjectID holds the string denoting the subject_id field in the database.
	FieldSubjectID = "subject_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSubject holds the string denoting the subject edge name in mutations.
	EdgeSubject = "subject"
	// EdgeSharedGroup holds the string denoting the shared_group edge name in mutations.
	EdgeSharedGroup = "shared_group"
	// Table holds the table name of the timer in the database.
	Table = "timers"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "timers"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// SubjectTable is the table that holds the subject relation/edge.
	SubjectTable = "timers"
	// SubjectInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectInverseTable = "subjects"
	// SubjectColumn is the table column denoting the subject relation/edge.
	SubjectColumn = "subject_id"
	// SharedGroupTable is the table that holds the shared_group relation/edge. The primary key declared below.
	SharedGroupTable = "timer_shared_group"
	// SharedGroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	SharedGroupInverseTable = "groups"
)

// Columns holds all SQL columns for timer fields.
var Columns = []string{
	FieldID,
	FieldStartAt,
	FieldContent,
	FieldSubjectID,
	FieldUserID,
}

var (
	// SharedGroupPrimaryKey and SharedGroupColumn2 are the table columns denoting the
	// primary key for the shared_group relation (M2M).
	SharedGroupPrimaryKey = []string{"timer_id", "group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Timer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStartAt orders the results by the start_at field.
func ByStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartAt, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// BySubjectID orders the results by the subject_id field.
func BySubjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubjectID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// BySubjectField orders the results by subject field.
func BySubjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectStep(), sql.OrderByField(field, opts...))
	}
}

// BySharedGroupCount orders the results by shared_group count.
func BySharedGroupCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSharedGroupStep(), opts...)
	}
}

// BySharedGroup orders the results by shared_group terms.
func BySharedGroup(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSharedGroupStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newSubjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
	)
}
func newSharedGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SharedGroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SharedGroupTable, SharedGroupPrimaryKey...),
	)
}
