// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AffiliationsColumns holds the columns for the "affiliations" table.
	AffiliationsColumns = []*schema.Column{
		{Name: "nickname", Type: field.TypeString},
		{Name: "role", Type: field.TypeInt8},
		{Name: "joined_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// AffiliationsTable holds the schema information for the "affiliations" table.
	AffiliationsTable = &schema.Table{
		Name:       "affiliations",
		Columns:    AffiliationsColumns,
		PrimaryKey: []*schema.Column{AffiliationsColumns[3], AffiliationsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "affiliations_users_user",
				Columns:    []*schema.Column{AffiliationsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "affiliations_groups_group",
				Columns:    []*schema.Column{AffiliationsColumns[4]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "affiliation_user_id",
				Unique:  false,
				Columns: []*schema.Column{AffiliationsColumns[3]},
			},
			{
				Name:    "affiliation_group_id",
				Unique:  false,
				Columns: []*schema.Column{AffiliationsColumns[4]},
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_owned_categories", Type: field.TypeUUID},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_users_owned_categories",
				Columns:    []*schema.Column{CategoriesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "category_user_owned_categories",
				Unique:  false,
				Columns: []*schema.Column{CategoriesColumns[2]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "nickname_policy", Type: field.TypeString, Default: ""},
		{Name: "reveal_policy", Type: field.TypeInt8},
		{Name: "invite_policy", Type: field.TypeInt8},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_owned_groups", Type: field.TypeUUID},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_users_owned_groups",
				Columns:    []*schema.Column{GroupsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// InviteCodesColumns holds the columns for the "invite_codes" table.
	InviteCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "used", Type: field.TypeInt32, Default: 0},
		{Name: "group_invite_codes", Type: field.TypeUUID},
	}
	// InviteCodesTable holds the schema information for the "invite_codes" table.
	InviteCodesTable = &schema.Table{
		Name:       "invite_codes",
		Columns:    InviteCodesColumns,
		PrimaryKey: []*schema.Column{InviteCodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "invite_codes_groups_invite_codes",
				Columns:    []*schema.Column{InviteCodesColumns[3]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_refresh_tokens", Type: field.TypeUUID},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refresh_tokens_users_refresh_tokens",
				Columns:    []*schema.Column{RefreshTokensColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StudyLogsColumns holds the columns for the "study_logs" table.
	StudyLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "subject_study_logs", Type: field.TypeUUID},
		{Name: "user_study_logs", Type: field.TypeUUID},
	}
	// StudyLogsTable holds the schema information for the "study_logs" table.
	StudyLogsTable = &schema.Table{
		Name:       "study_logs",
		Columns:    StudyLogsColumns,
		PrimaryKey: []*schema.Column{StudyLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "study_logs_subjects_study_logs",
				Columns:    []*schema.Column{StudyLogsColumns[4]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "study_logs_users_study_logs",
				Columns:    []*schema.Column{StudyLogsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "studylog_user_study_logs",
				Unique:  false,
				Columns: []*schema.Column{StudyLogsColumns[5]},
			},
			{
				Name:    "studylog_user_study_logs_subject_study_logs",
				Unique:  false,
				Columns: []*schema.Column{StudyLogsColumns[5], StudyLogsColumns[4]},
			},
			{
				Name:    "studylog_start_at",
				Unique:  false,
				Columns: []*schema.Column{StudyLogsColumns[1]},
			},
			{
				Name:    "studylog_end_at",
				Unique:  false,
				Columns: []*schema.Column{StudyLogsColumns[2]},
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "color", Type: field.TypeString},
		{Name: "category_subjects", Type: field.TypeUUID},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subjects_categories_subjects",
				Columns:    []*schema.Column{SubjectsColumns[3]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subject_category_subjects",
				Unique:  false,
				Columns: []*schema.Column{SubjectsColumns[3]},
			},
		},
	}
	// TimersColumns holds the columns for the "timers" table.
	TimersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "start_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "subject_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
	}
	// TimersTable holds the schema information for the "timers" table.
	TimersTable = &schema.Table{
		Name:       "timers",
		Columns:    TimersColumns,
		PrimaryKey: []*schema.Column{TimersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timers_subjects_timers",
				Columns:    []*schema.Column{TimersColumns[3]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "timers_users_timers",
				Columns:    []*schema.Column{TimersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Default: "유저"},
		{Name: "oauth_id", Type: field.TypeString},
		{Name: "oauth_provider", Type: field.TypeInt8},
		{Name: "terms_agreed", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_oauth_id_oauth_provider",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2], UsersColumns[3]},
			},
		},
	}
	// StudyLogSharedGroupColumns holds the columns for the "study_log_shared_group" table.
	StudyLogSharedGroupColumns = []*schema.Column{
		{Name: "study_log_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// StudyLogSharedGroupTable holds the schema information for the "study_log_shared_group" table.
	StudyLogSharedGroupTable = &schema.Table{
		Name:       "study_log_shared_group",
		Columns:    StudyLogSharedGroupColumns,
		PrimaryKey: []*schema.Column{StudyLogSharedGroupColumns[0], StudyLogSharedGroupColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "study_log_shared_group_study_log_id",
				Columns:    []*schema.Column{StudyLogSharedGroupColumns[0]},
				RefColumns: []*schema.Column{StudyLogsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "study_log_shared_group_group_id",
				Columns:    []*schema.Column{StudyLogSharedGroupColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TimerSharedGroupColumns holds the columns for the "timer_shared_group" table.
	TimerSharedGroupColumns = []*schema.Column{
		{Name: "timer_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// TimerSharedGroupTable holds the schema information for the "timer_shared_group" table.
	TimerSharedGroupTable = &schema.Table{
		Name:       "timer_shared_group",
		Columns:    TimerSharedGroupColumns,
		PrimaryKey: []*schema.Column{TimerSharedGroupColumns[0], TimerSharedGroupColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timer_shared_group_timer_id",
				Columns:    []*schema.Column{TimerSharedGroupColumns[0]},
				RefColumns: []*schema.Column{TimersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "timer_shared_group_group_id",
				Columns:    []*schema.Column{TimerSharedGroupColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AffiliationsTable,
		CategoriesTable,
		GroupsTable,
		InviteCodesTable,
		RefreshTokensTable,
		StudyLogsTable,
		SubjectsTable,
		TimersTable,
		UsersTable,
		StudyLogSharedGroupTable,
		TimerSharedGroupTable,
	}
)

func init() {
	AffiliationsTable.ForeignKeys[0].RefTable = UsersTable
	AffiliationsTable.ForeignKeys[1].RefTable = GroupsTable
	CategoriesTable.ForeignKeys[0].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = UsersTable
	InviteCodesTable.ForeignKeys[0].RefTable = GroupsTable
	RefreshTokensTable.ForeignKeys[0].RefTable = UsersTable
	StudyLogsTable.ForeignKeys[0].RefTable = SubjectsTable
	StudyLogsTable.ForeignKeys[1].RefTable = UsersTable
	SubjectsTable.ForeignKeys[0].RefTable = CategoriesTable
	TimersTable.ForeignKeys[0].RefTable = SubjectsTable
	TimersTable.ForeignKeys[1].RefTable = UsersTable
	StudyLogSharedGroupTable.ForeignKeys[0].RefTable = StudyLogsTable
	StudyLogSharedGroupTable.ForeignKeys[1].RefTable = GroupsTable
	TimerSharedGroupTable.ForeignKeys[0].RefTable = TimersTable
	TimerSharedGroupTable.ForeignKeys[1].RefTable = GroupsTable
}
