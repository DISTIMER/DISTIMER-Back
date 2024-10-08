// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/predicate"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
)

// StudyLogUpdate is the builder for updating StudyLog entities.
type StudyLogUpdate struct {
	config
	hooks    []Hook
	mutation *StudyLogMutation
}

// Where appends a list predicates to the StudyLogUpdate builder.
func (slu *StudyLogUpdate) Where(ps ...predicate.StudyLog) *StudyLogUpdate {
	slu.mutation.Where(ps...)
	return slu
}

// SetStartAt sets the "start_at" field.
func (slu *StudyLogUpdate) SetStartAt(t time.Time) *StudyLogUpdate {
	slu.mutation.SetStartAt(t)
	return slu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (slu *StudyLogUpdate) SetNillableStartAt(t *time.Time) *StudyLogUpdate {
	if t != nil {
		slu.SetStartAt(*t)
	}
	return slu
}

// SetEndAt sets the "end_at" field.
func (slu *StudyLogUpdate) SetEndAt(t time.Time) *StudyLogUpdate {
	slu.mutation.SetEndAt(t)
	return slu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (slu *StudyLogUpdate) SetNillableEndAt(t *time.Time) *StudyLogUpdate {
	if t != nil {
		slu.SetEndAt(*t)
	}
	return slu
}

// SetContent sets the "content" field.
func (slu *StudyLogUpdate) SetContent(s string) *StudyLogUpdate {
	slu.mutation.SetContent(s)
	return slu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (slu *StudyLogUpdate) SetNillableContent(s *string) *StudyLogUpdate {
	if s != nil {
		slu.SetContent(*s)
	}
	return slu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (slu *StudyLogUpdate) SetUserID(id uuid.UUID) *StudyLogUpdate {
	slu.mutation.SetUserID(id)
	return slu
}

// SetUser sets the "user" edge to the User entity.
func (slu *StudyLogUpdate) SetUser(u *User) *StudyLogUpdate {
	return slu.SetUserID(u.ID)
}

// SetSubjectID sets the "subject" edge to the Subject entity by ID.
func (slu *StudyLogUpdate) SetSubjectID(id uuid.UUID) *StudyLogUpdate {
	slu.mutation.SetSubjectID(id)
	return slu
}

// SetSubject sets the "subject" edge to the Subject entity.
func (slu *StudyLogUpdate) SetSubject(s *Subject) *StudyLogUpdate {
	return slu.SetSubjectID(s.ID)
}

// AddSharedGroupIDs adds the "shared_group" edge to the Group entity by IDs.
func (slu *StudyLogUpdate) AddSharedGroupIDs(ids ...uuid.UUID) *StudyLogUpdate {
	slu.mutation.AddSharedGroupIDs(ids...)
	return slu
}

// AddSharedGroup adds the "shared_group" edges to the Group entity.
func (slu *StudyLogUpdate) AddSharedGroup(g ...*Group) *StudyLogUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return slu.AddSharedGroupIDs(ids...)
}

// Mutation returns the StudyLogMutation object of the builder.
func (slu *StudyLogUpdate) Mutation() *StudyLogMutation {
	return slu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (slu *StudyLogUpdate) ClearUser() *StudyLogUpdate {
	slu.mutation.ClearUser()
	return slu
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (slu *StudyLogUpdate) ClearSubject() *StudyLogUpdate {
	slu.mutation.ClearSubject()
	return slu
}

// ClearSharedGroup clears all "shared_group" edges to the Group entity.
func (slu *StudyLogUpdate) ClearSharedGroup() *StudyLogUpdate {
	slu.mutation.ClearSharedGroup()
	return slu
}

// RemoveSharedGroupIDs removes the "shared_group" edge to Group entities by IDs.
func (slu *StudyLogUpdate) RemoveSharedGroupIDs(ids ...uuid.UUID) *StudyLogUpdate {
	slu.mutation.RemoveSharedGroupIDs(ids...)
	return slu
}

// RemoveSharedGroup removes "shared_group" edges to Group entities.
func (slu *StudyLogUpdate) RemoveSharedGroup(g ...*Group) *StudyLogUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return slu.RemoveSharedGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *StudyLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, slu.sqlSave, slu.mutation, slu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (slu *StudyLogUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *StudyLogUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *StudyLogUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slu *StudyLogUpdate) check() error {
	if _, ok := slu.mutation.UserID(); slu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "StudyLog.user"`)
	}
	if _, ok := slu.mutation.SubjectID(); slu.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "StudyLog.subject"`)
	}
	return nil
}

func (slu *StudyLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := slu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(studylog.Table, studylog.Columns, sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID))
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := slu.mutation.StartAt(); ok {
		_spec.SetField(studylog.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := slu.mutation.EndAt(); ok {
		_spec.SetField(studylog.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := slu.mutation.Content(); ok {
		_spec.SetField(studylog.FieldContent, field.TypeString, value)
	}
	if slu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.UserTable,
			Columns: []string{studylog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := slu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.UserTable,
			Columns: []string{studylog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if slu.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.SubjectTable,
			Columns: []string{studylog.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := slu.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.SubjectTable,
			Columns: []string{studylog.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if slu.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   studylog.SharedGroupTable,
			Columns: studylog.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := slu.mutation.RemovedSharedGroupIDs(); len(nodes) > 0 && !slu.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   studylog.SharedGroupTable,
			Columns: studylog.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := slu.mutation.SharedGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   studylog.SharedGroupTable,
			Columns: studylog.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studylog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	slu.mutation.done = true
	return n, nil
}

// StudyLogUpdateOne is the builder for updating a single StudyLog entity.
type StudyLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudyLogMutation
}

// SetStartAt sets the "start_at" field.
func (sluo *StudyLogUpdateOne) SetStartAt(t time.Time) *StudyLogUpdateOne {
	sluo.mutation.SetStartAt(t)
	return sluo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (sluo *StudyLogUpdateOne) SetNillableStartAt(t *time.Time) *StudyLogUpdateOne {
	if t != nil {
		sluo.SetStartAt(*t)
	}
	return sluo
}

// SetEndAt sets the "end_at" field.
func (sluo *StudyLogUpdateOne) SetEndAt(t time.Time) *StudyLogUpdateOne {
	sluo.mutation.SetEndAt(t)
	return sluo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (sluo *StudyLogUpdateOne) SetNillableEndAt(t *time.Time) *StudyLogUpdateOne {
	if t != nil {
		sluo.SetEndAt(*t)
	}
	return sluo
}

// SetContent sets the "content" field.
func (sluo *StudyLogUpdateOne) SetContent(s string) *StudyLogUpdateOne {
	sluo.mutation.SetContent(s)
	return sluo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (sluo *StudyLogUpdateOne) SetNillableContent(s *string) *StudyLogUpdateOne {
	if s != nil {
		sluo.SetContent(*s)
	}
	return sluo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sluo *StudyLogUpdateOne) SetUserID(id uuid.UUID) *StudyLogUpdateOne {
	sluo.mutation.SetUserID(id)
	return sluo
}

// SetUser sets the "user" edge to the User entity.
func (sluo *StudyLogUpdateOne) SetUser(u *User) *StudyLogUpdateOne {
	return sluo.SetUserID(u.ID)
}

// SetSubjectID sets the "subject" edge to the Subject entity by ID.
func (sluo *StudyLogUpdateOne) SetSubjectID(id uuid.UUID) *StudyLogUpdateOne {
	sluo.mutation.SetSubjectID(id)
	return sluo
}

// SetSubject sets the "subject" edge to the Subject entity.
func (sluo *StudyLogUpdateOne) SetSubject(s *Subject) *StudyLogUpdateOne {
	return sluo.SetSubjectID(s.ID)
}

// AddSharedGroupIDs adds the "shared_group" edge to the Group entity by IDs.
func (sluo *StudyLogUpdateOne) AddSharedGroupIDs(ids ...uuid.UUID) *StudyLogUpdateOne {
	sluo.mutation.AddSharedGroupIDs(ids...)
	return sluo
}

// AddSharedGroup adds the "shared_group" edges to the Group entity.
func (sluo *StudyLogUpdateOne) AddSharedGroup(g ...*Group) *StudyLogUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return sluo.AddSharedGroupIDs(ids...)
}

// Mutation returns the StudyLogMutation object of the builder.
func (sluo *StudyLogUpdateOne) Mutation() *StudyLogMutation {
	return sluo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (sluo *StudyLogUpdateOne) ClearUser() *StudyLogUpdateOne {
	sluo.mutation.ClearUser()
	return sluo
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (sluo *StudyLogUpdateOne) ClearSubject() *StudyLogUpdateOne {
	sluo.mutation.ClearSubject()
	return sluo
}

// ClearSharedGroup clears all "shared_group" edges to the Group entity.
func (sluo *StudyLogUpdateOne) ClearSharedGroup() *StudyLogUpdateOne {
	sluo.mutation.ClearSharedGroup()
	return sluo
}

// RemoveSharedGroupIDs removes the "shared_group" edge to Group entities by IDs.
func (sluo *StudyLogUpdateOne) RemoveSharedGroupIDs(ids ...uuid.UUID) *StudyLogUpdateOne {
	sluo.mutation.RemoveSharedGroupIDs(ids...)
	return sluo
}

// RemoveSharedGroup removes "shared_group" edges to Group entities.
func (sluo *StudyLogUpdateOne) RemoveSharedGroup(g ...*Group) *StudyLogUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return sluo.RemoveSharedGroupIDs(ids...)
}

// Where appends a list predicates to the StudyLogUpdate builder.
func (sluo *StudyLogUpdateOne) Where(ps ...predicate.StudyLog) *StudyLogUpdateOne {
	sluo.mutation.Where(ps...)
	return sluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *StudyLogUpdateOne) Select(field string, fields ...string) *StudyLogUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated StudyLog entity.
func (sluo *StudyLogUpdateOne) Save(ctx context.Context) (*StudyLog, error) {
	return withHooks(ctx, sluo.sqlSave, sluo.mutation, sluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *StudyLogUpdateOne) SaveX(ctx context.Context) *StudyLog {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *StudyLogUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *StudyLogUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sluo *StudyLogUpdateOne) check() error {
	if _, ok := sluo.mutation.UserID(); sluo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "StudyLog.user"`)
	}
	if _, ok := sluo.mutation.SubjectID(); sluo.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "StudyLog.subject"`)
	}
	return nil
}

func (sluo *StudyLogUpdateOne) sqlSave(ctx context.Context) (_node *StudyLog, err error) {
	if err := sluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(studylog.Table, studylog.Columns, sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID))
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StudyLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studylog.FieldID)
		for _, f := range fields {
			if !studylog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != studylog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sluo.mutation.StartAt(); ok {
		_spec.SetField(studylog.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := sluo.mutation.EndAt(); ok {
		_spec.SetField(studylog.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := sluo.mutation.Content(); ok {
		_spec.SetField(studylog.FieldContent, field.TypeString, value)
	}
	if sluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.UserTable,
			Columns: []string{studylog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.UserTable,
			Columns: []string{studylog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sluo.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.SubjectTable,
			Columns: []string{studylog.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sluo.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studylog.SubjectTable,
			Columns: []string{studylog.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sluo.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   studylog.SharedGroupTable,
			Columns: studylog.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sluo.mutation.RemovedSharedGroupIDs(); len(nodes) > 0 && !sluo.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   studylog.SharedGroupTable,
			Columns: studylog.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sluo.mutation.SharedGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   studylog.SharedGroupTable,
			Columns: studylog.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudyLog{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studylog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sluo.mutation.done = true
	return _node, nil
}
