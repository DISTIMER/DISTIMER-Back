// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOauthID holds the string denoting the oauth_id field in the database.
	FieldOauthID = "oauth_id"
	// FieldOauthProvider holds the string denoting the oauth_provider field in the database.
	FieldOauthProvider = "oauth_provider"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeJoinedGroups holds the string denoting the joined_groups edge name in mutations.
	EdgeJoinedGroups = "joined_groups"
	// EdgeOwnedGroups holds the string denoting the owned_groups edge name in mutations.
	EdgeOwnedGroups = "owned_groups"
	// EdgeStudyLogs holds the string denoting the study_logs edge name in mutations.
	EdgeStudyLogs = "study_logs"
	// EdgeRefreshTokens holds the string denoting the refresh_tokens edge name in mutations.
	EdgeRefreshTokens = "refresh_tokens"
	// EdgeAffiliations holds the string denoting the affiliations edge name in mutations.
	EdgeAffiliations = "affiliations"
	// Table holds the table name of the user in the database.
	Table = "users"
	// JoinedGroupsTable is the table that holds the joined_groups relation/edge. The primary key declared below.
	JoinedGroupsTable = "affiliations"
	// JoinedGroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	JoinedGroupsInverseTable = "groups"
	// OwnedGroupsTable is the table that holds the owned_groups relation/edge.
	OwnedGroupsTable = "groups"
	// OwnedGroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	OwnedGroupsInverseTable = "groups"
	// OwnedGroupsColumn is the table column denoting the owned_groups relation/edge.
	OwnedGroupsColumn = "user_owned_groups"
	// StudyLogsTable is the table that holds the study_logs relation/edge.
	StudyLogsTable = "study_logs"
	// StudyLogsInverseTable is the table name for the StudyLog entity.
	// It exists in this package in order to avoid circular dependency with the "studylog" package.
	StudyLogsInverseTable = "study_logs"
	// StudyLogsColumn is the table column denoting the study_logs relation/edge.
	StudyLogsColumn = "user_study_logs"
	// RefreshTokensTable is the table that holds the refresh_tokens relation/edge.
	RefreshTokensTable = "refresh_tokens"
	// RefreshTokensInverseTable is the table name for the RefreshToken entity.
	// It exists in this package in order to avoid circular dependency with the "refreshtoken" package.
	RefreshTokensInverseTable = "refresh_tokens"
	// RefreshTokensColumn is the table column denoting the refresh_tokens relation/edge.
	RefreshTokensColumn = "user_refresh_tokens"
	// AffiliationsTable is the table that holds the affiliations relation/edge.
	AffiliationsTable = "affiliations"
	// AffiliationsInverseTable is the table name for the Affiliation entity.
	// It exists in this package in order to avoid circular dependency with the "affiliation" package.
	AffiliationsInverseTable = "affiliations"
	// AffiliationsColumn is the table column denoting the affiliations relation/edge.
	AffiliationsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldOauthID,
	FieldOauthProvider,
	FieldCreatedAt,
}

var (
	// JoinedGroupsPrimaryKey and JoinedGroupsColumn2 are the table columns denoting the
	// primary key for the joined_groups relation (M2M).
	JoinedGroupsPrimaryKey = []string{"user_id", "group_id"}
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOauthID orders the results by the oauth_id field.
func ByOauthID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOauthID, opts...).ToFunc()
}

// ByOauthProvider orders the results by the oauth_provider field.
func ByOauthProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOauthProvider, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByJoinedGroupsCount orders the results by joined_groups count.
func ByJoinedGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newJoinedGroupsStep(), opts...)
	}
}

// ByJoinedGroups orders the results by joined_groups terms.
func ByJoinedGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJoinedGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnedGroupsCount orders the results by owned_groups count.
func ByOwnedGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOwnedGroupsStep(), opts...)
	}
}

// ByOwnedGroups orders the results by owned_groups terms.
func ByOwnedGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnedGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByRefreshTokensCount orders the results by refresh_tokens count.
func ByRefreshTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRefreshTokensStep(), opts...)
	}
}

// ByRefreshTokens orders the results by refresh_tokens terms.
func ByRefreshTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRefreshTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAffiliationsCount orders the results by affiliations count.
func ByAffiliationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAffiliationsStep(), opts...)
	}
}

// ByAffiliations orders the results by affiliations terms.
func ByAffiliations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAffiliationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newJoinedGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JoinedGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, JoinedGroupsTable, JoinedGroupsPrimaryKey...),
	)
}
func newOwnedGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnedGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OwnedGroupsTable, OwnedGroupsColumn),
	)
}
func newStudyLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudyLogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudyLogsTable, StudyLogsColumn),
	)
}
func newRefreshTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RefreshTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RefreshTokensTable, RefreshTokensColumn),
	)
}
func newAffiliationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AffiliationsInverseTable, AffiliationsColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, AffiliationsTable, AffiliationsColumn),
	)
}
