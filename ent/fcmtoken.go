// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/fcmtoken"
	"pentag.kr/distimer/ent/session"
)

// FCMToken is the model entity for the FCMToken schema.
type FCMToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PushToken holds the value of the "push_token" field.
	PushToken string `json:"push_token,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FCMTokenQuery when eager-loading is set.
	Edges             FCMTokenEdges `json:"edges"`
	session_fcm_token *uuid.UUID
	selectValues      sql.SelectValues
}

// FCMTokenEdges holds the relations/edges for other nodes in the graph.
type FCMTokenEdges struct {
	// Session holds the value of the session edge.
	Session *Session `json:"session,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SessionOrErr returns the Session value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FCMTokenEdges) SessionOrErr() (*Session, error) {
	if e.Session != nil {
		return e.Session, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: session.Label}
	}
	return nil, &NotLoadedError{edge: "session"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FCMToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fcmtoken.FieldID:
			values[i] = new(sql.NullInt64)
		case fcmtoken.FieldPushToken:
			values[i] = new(sql.NullString)
		case fcmtoken.ForeignKeys[0]: // session_fcm_token
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FCMToken fields.
func (ft *FCMToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fcmtoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ft.ID = int(value.Int64)
		case fcmtoken.FieldPushToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field push_token", values[i])
			} else if value.Valid {
				ft.PushToken = value.String
			}
		case fcmtoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field session_fcm_token", values[i])
			} else if value.Valid {
				ft.session_fcm_token = new(uuid.UUID)
				*ft.session_fcm_token = *value.S.(*uuid.UUID)
			}
		default:
			ft.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FCMToken.
// This includes values selected through modifiers, order, etc.
func (ft *FCMToken) Value(name string) (ent.Value, error) {
	return ft.selectValues.Get(name)
}

// QuerySession queries the "session" edge of the FCMToken entity.
func (ft *FCMToken) QuerySession() *SessionQuery {
	return NewFCMTokenClient(ft.config).QuerySession(ft)
}

// Update returns a builder for updating this FCMToken.
// Note that you need to call FCMToken.Unwrap() before calling this method if this FCMToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FCMToken) Update() *FCMTokenUpdateOne {
	return NewFCMTokenClient(ft.config).UpdateOne(ft)
}

// Unwrap unwraps the FCMToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FCMToken) Unwrap() *FCMToken {
	_tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FCMToken is not a transactional entity")
	}
	ft.config.driver = _tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FCMToken) String() string {
	var builder strings.Builder
	builder.WriteString("FCMToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ft.ID))
	builder.WriteString("push_token=")
	builder.WriteString(ft.PushToken)
	builder.WriteByte(')')
	return builder.String()
}

// FCMTokens is a parsable slice of FCMToken.
type FCMTokens []*FCMToken
