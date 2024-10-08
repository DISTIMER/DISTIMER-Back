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
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
)

// TimerUpdate is the builder for updating Timer entities.
type TimerUpdate struct {
	config
	hooks    []Hook
	mutation *TimerMutation
}

// Where appends a list predicates to the TimerUpdate builder.
func (tu *TimerUpdate) Where(ps ...predicate.Timer) *TimerUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetStartAt sets the "start_at" field.
func (tu *TimerUpdate) SetStartAt(t time.Time) *TimerUpdate {
	tu.mutation.SetStartAt(t)
	return tu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (tu *TimerUpdate) SetNillableStartAt(t *time.Time) *TimerUpdate {
	if t != nil {
		tu.SetStartAt(*t)
	}
	return tu
}

// SetContent sets the "content" field.
func (tu *TimerUpdate) SetContent(s string) *TimerUpdate {
	tu.mutation.SetContent(s)
	return tu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (tu *TimerUpdate) SetNillableContent(s *string) *TimerUpdate {
	if s != nil {
		tu.SetContent(*s)
	}
	return tu
}

// SetSubjectID sets the "subject_id" field.
func (tu *TimerUpdate) SetSubjectID(u uuid.UUID) *TimerUpdate {
	tu.mutation.SetSubjectID(u)
	return tu
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (tu *TimerUpdate) SetNillableSubjectID(u *uuid.UUID) *TimerUpdate {
	if u != nil {
		tu.SetSubjectID(*u)
	}
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TimerUpdate) SetUserID(u uuid.UUID) *TimerUpdate {
	tu.mutation.SetUserID(u)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TimerUpdate) SetNillableUserID(u *uuid.UUID) *TimerUpdate {
	if u != nil {
		tu.SetUserID(*u)
	}
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TimerUpdate) SetUser(u *User) *TimerUpdate {
	return tu.SetUserID(u.ID)
}

// SetSubject sets the "subject" edge to the Subject entity.
func (tu *TimerUpdate) SetSubject(s *Subject) *TimerUpdate {
	return tu.SetSubjectID(s.ID)
}

// AddSharedGroupIDs adds the "shared_group" edge to the Group entity by IDs.
func (tu *TimerUpdate) AddSharedGroupIDs(ids ...uuid.UUID) *TimerUpdate {
	tu.mutation.AddSharedGroupIDs(ids...)
	return tu
}

// AddSharedGroup adds the "shared_group" edges to the Group entity.
func (tu *TimerUpdate) AddSharedGroup(g ...*Group) *TimerUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddSharedGroupIDs(ids...)
}

// Mutation returns the TimerMutation object of the builder.
func (tu *TimerUpdate) Mutation() *TimerMutation {
	return tu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TimerUpdate) ClearUser() *TimerUpdate {
	tu.mutation.ClearUser()
	return tu
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (tu *TimerUpdate) ClearSubject() *TimerUpdate {
	tu.mutation.ClearSubject()
	return tu
}

// ClearSharedGroup clears all "shared_group" edges to the Group entity.
func (tu *TimerUpdate) ClearSharedGroup() *TimerUpdate {
	tu.mutation.ClearSharedGroup()
	return tu
}

// RemoveSharedGroupIDs removes the "shared_group" edge to Group entities by IDs.
func (tu *TimerUpdate) RemoveSharedGroupIDs(ids ...uuid.UUID) *TimerUpdate {
	tu.mutation.RemoveSharedGroupIDs(ids...)
	return tu
}

// RemoveSharedGroup removes "shared_group" edges to Group entities.
func (tu *TimerUpdate) RemoveSharedGroup(g ...*Group) *TimerUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveSharedGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TimerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TimerUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TimerUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TimerUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TimerUpdate) check() error {
	if _, ok := tu.mutation.UserID(); tu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Timer.user"`)
	}
	if _, ok := tu.mutation.SubjectID(); tu.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Timer.subject"`)
	}
	return nil
}

func (tu *TimerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(timer.Table, timer.Columns, sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.StartAt(); ok {
		_spec.SetField(timer.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Content(); ok {
		_spec.SetField(timer.FieldContent, field.TypeString, value)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   timer.UserTable,
			Columns: []string{timer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   timer.UserTable,
			Columns: []string{timer.UserColumn},
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
	if tu.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timer.SubjectTable,
			Columns: []string{timer.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timer.SubjectTable,
			Columns: []string{timer.SubjectColumn},
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
	if tu.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSharedGroupIDs(); len(nodes) > 0 && !tu.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
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
	if nodes := tu.mutation.SharedGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TimerUpdateOne is the builder for updating a single Timer entity.
type TimerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TimerMutation
}

// SetStartAt sets the "start_at" field.
func (tuo *TimerUpdateOne) SetStartAt(t time.Time) *TimerUpdateOne {
	tuo.mutation.SetStartAt(t)
	return tuo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (tuo *TimerUpdateOne) SetNillableStartAt(t *time.Time) *TimerUpdateOne {
	if t != nil {
		tuo.SetStartAt(*t)
	}
	return tuo
}

// SetContent sets the "content" field.
func (tuo *TimerUpdateOne) SetContent(s string) *TimerUpdateOne {
	tuo.mutation.SetContent(s)
	return tuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (tuo *TimerUpdateOne) SetNillableContent(s *string) *TimerUpdateOne {
	if s != nil {
		tuo.SetContent(*s)
	}
	return tuo
}

// SetSubjectID sets the "subject_id" field.
func (tuo *TimerUpdateOne) SetSubjectID(u uuid.UUID) *TimerUpdateOne {
	tuo.mutation.SetSubjectID(u)
	return tuo
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (tuo *TimerUpdateOne) SetNillableSubjectID(u *uuid.UUID) *TimerUpdateOne {
	if u != nil {
		tuo.SetSubjectID(*u)
	}
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TimerUpdateOne) SetUserID(u uuid.UUID) *TimerUpdateOne {
	tuo.mutation.SetUserID(u)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TimerUpdateOne) SetNillableUserID(u *uuid.UUID) *TimerUpdateOne {
	if u != nil {
		tuo.SetUserID(*u)
	}
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TimerUpdateOne) SetUser(u *User) *TimerUpdateOne {
	return tuo.SetUserID(u.ID)
}

// SetSubject sets the "subject" edge to the Subject entity.
func (tuo *TimerUpdateOne) SetSubject(s *Subject) *TimerUpdateOne {
	return tuo.SetSubjectID(s.ID)
}

// AddSharedGroupIDs adds the "shared_group" edge to the Group entity by IDs.
func (tuo *TimerUpdateOne) AddSharedGroupIDs(ids ...uuid.UUID) *TimerUpdateOne {
	tuo.mutation.AddSharedGroupIDs(ids...)
	return tuo
}

// AddSharedGroup adds the "shared_group" edges to the Group entity.
func (tuo *TimerUpdateOne) AddSharedGroup(g ...*Group) *TimerUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddSharedGroupIDs(ids...)
}

// Mutation returns the TimerMutation object of the builder.
func (tuo *TimerUpdateOne) Mutation() *TimerMutation {
	return tuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TimerUpdateOne) ClearUser() *TimerUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (tuo *TimerUpdateOne) ClearSubject() *TimerUpdateOne {
	tuo.mutation.ClearSubject()
	return tuo
}

// ClearSharedGroup clears all "shared_group" edges to the Group entity.
func (tuo *TimerUpdateOne) ClearSharedGroup() *TimerUpdateOne {
	tuo.mutation.ClearSharedGroup()
	return tuo
}

// RemoveSharedGroupIDs removes the "shared_group" edge to Group entities by IDs.
func (tuo *TimerUpdateOne) RemoveSharedGroupIDs(ids ...uuid.UUID) *TimerUpdateOne {
	tuo.mutation.RemoveSharedGroupIDs(ids...)
	return tuo
}

// RemoveSharedGroup removes "shared_group" edges to Group entities.
func (tuo *TimerUpdateOne) RemoveSharedGroup(g ...*Group) *TimerUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveSharedGroupIDs(ids...)
}

// Where appends a list predicates to the TimerUpdate builder.
func (tuo *TimerUpdateOne) Where(ps ...predicate.Timer) *TimerUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TimerUpdateOne) Select(field string, fields ...string) *TimerUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Timer entity.
func (tuo *TimerUpdateOne) Save(ctx context.Context) (*Timer, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TimerUpdateOne) SaveX(ctx context.Context) *Timer {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TimerUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TimerUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TimerUpdateOne) check() error {
	if _, ok := tuo.mutation.UserID(); tuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Timer.user"`)
	}
	if _, ok := tuo.mutation.SubjectID(); tuo.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Timer.subject"`)
	}
	return nil
}

func (tuo *TimerUpdateOne) sqlSave(ctx context.Context) (_node *Timer, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(timer.Table, timer.Columns, sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Timer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timer.FieldID)
		for _, f := range fields {
			if !timer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != timer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.StartAt(); ok {
		_spec.SetField(timer.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Content(); ok {
		_spec.SetField(timer.FieldContent, field.TypeString, value)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   timer.UserTable,
			Columns: []string{timer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   timer.UserTable,
			Columns: []string{timer.UserColumn},
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
	if tuo.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timer.SubjectTable,
			Columns: []string{timer.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timer.SubjectTable,
			Columns: []string{timer.SubjectColumn},
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
	if tuo.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSharedGroupIDs(); len(nodes) > 0 && !tuo.mutation.SharedGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
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
	if nodes := tuo.mutation.SharedGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
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
	_node = &Timer{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
