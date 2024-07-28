// Code generated by ent, DO NOT EDIT.

package studylog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLTE(FieldID, id))
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldStartAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldEndAt, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldContent, v))
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldStartAt, v))
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNEQ(FieldStartAt, v))
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldIn(FieldStartAt, vs...))
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNotIn(FieldStartAt, vs...))
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGT(FieldStartAt, v))
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGTE(FieldStartAt, v))
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLT(FieldStartAt, v))
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLTE(FieldStartAt, v))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLTE(FieldEndAt, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.StudyLog {
	return predicate.StudyLog(sql.FieldContainsFold(FieldContent, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.StudyLog {
	return predicate.StudyLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.StudyLog {
	return predicate.StudyLog(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.StudyLog {
	return predicate.StudyLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.StudyLog {
	return predicate.StudyLog(func(s *sql.Selector) {
		step := newSubjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSharedGroup applies the HasEdge predicate on the "shared_group" edge.
func HasSharedGroup() predicate.StudyLog {
	return predicate.StudyLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SharedGroupTable, SharedGroupPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSharedGroupWith applies the HasEdge predicate on the "shared_group" edge with a given conditions (other predicates).
func HasSharedGroupWith(preds ...predicate.Group) predicate.StudyLog {
	return predicate.StudyLog(func(s *sql.Selector) {
		step := newSharedGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StudyLog) predicate.StudyLog {
	return predicate.StudyLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StudyLog) predicate.StudyLog {
	return predicate.StudyLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StudyLog) predicate.StudyLog {
	return predicate.StudyLog(sql.NotPredicates(p))
}
