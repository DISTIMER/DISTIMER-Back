// Code generated by ent, DO NOT EDIT.

package subject

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subject type in the database.
	Label = "subject"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeStudyLogs holds the string denoting the study_logs edge name in mutations.
	EdgeStudyLogs = "study_logs"
	// Table holds the table name of the subject in the database.
	Table = "subjects"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "subjects"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_subjects"
	// StudyLogsTable is the table that holds the study_logs relation/edge.
	StudyLogsTable = "study_logs"
	// StudyLogsInverseTable is the table name for the StudyLog entity.
	// It exists in this package in order to avoid circular dependency with the "studylog" package.
	StudyLogsInverseTable = "study_logs"
	// StudyLogsColumn is the table column denoting the study_logs relation/edge.
	StudyLogsColumn = "subject_study_logs"
)

// Columns holds all SQL columns for subject fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldColor,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subjects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"category_subjects",
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

// OrderOption defines the ordering options for the Subject queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByCategoryField orders the results by category field.
func ByCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByStudyLogsCount orders the results by study_logs count.
func ByStudyLogsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudyLogsStep(), opts...)
	}
}

// ByStudyLogs orders the results by study_logs terms.
func ByStudyLogs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudyLogsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
	)
}
func newStudyLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudyLogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudyLogsTable, StudyLogsColumn),
	)
}
