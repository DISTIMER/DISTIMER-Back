// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
)

// StudyLogCreate is the builder for creating a StudyLog entity.
type StudyLogCreate struct {
	config
	mutation *StudyLogMutation
	hooks    []Hook
}

// SetStartAt sets the "start_at" field.
func (slc *StudyLogCreate) SetStartAt(t time.Time) *StudyLogCreate {
	slc.mutation.SetStartAt(t)
	return slc
}

// SetEndAt sets the "end_at" field.
func (slc *StudyLogCreate) SetEndAt(t time.Time) *StudyLogCreate {
	slc.mutation.SetEndAt(t)
	return slc
}

// SetContent sets the "content" field.
func (slc *StudyLogCreate) SetContent(s string) *StudyLogCreate {
	slc.mutation.SetContent(s)
	return slc
}

// SetID sets the "id" field.
func (slc *StudyLogCreate) SetID(u uuid.UUID) *StudyLogCreate {
	slc.mutation.SetID(u)
	return slc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (slc *StudyLogCreate) SetNillableID(u *uuid.UUID) *StudyLogCreate {
	if u != nil {
		slc.SetID(*u)
	}
	return slc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (slc *StudyLogCreate) SetUserID(id uuid.UUID) *StudyLogCreate {
	slc.mutation.SetUserID(id)
	return slc
}

// SetUser sets the "user" edge to the User entity.
func (slc *StudyLogCreate) SetUser(u *User) *StudyLogCreate {
	return slc.SetUserID(u.ID)
}

// SetSubjectID sets the "subject" edge to the Subject entity by ID.
func (slc *StudyLogCreate) SetSubjectID(id uuid.UUID) *StudyLogCreate {
	slc.mutation.SetSubjectID(id)
	return slc
}

// SetSubject sets the "subject" edge to the Subject entity.
func (slc *StudyLogCreate) SetSubject(s *Subject) *StudyLogCreate {
	return slc.SetSubjectID(s.ID)
}

// AddSharedGroupIDs adds the "shared_group" edge to the Group entity by IDs.
func (slc *StudyLogCreate) AddSharedGroupIDs(ids ...uuid.UUID) *StudyLogCreate {
	slc.mutation.AddSharedGroupIDs(ids...)
	return slc
}

// AddSharedGroup adds the "shared_group" edges to the Group entity.
func (slc *StudyLogCreate) AddSharedGroup(g ...*Group) *StudyLogCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return slc.AddSharedGroupIDs(ids...)
}

// Mutation returns the StudyLogMutation object of the builder.
func (slc *StudyLogCreate) Mutation() *StudyLogMutation {
	return slc.mutation
}

// Save creates the StudyLog in the database.
func (slc *StudyLogCreate) Save(ctx context.Context) (*StudyLog, error) {
	slc.defaults()
	return withHooks(ctx, slc.sqlSave, slc.mutation, slc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (slc *StudyLogCreate) SaveX(ctx context.Context) *StudyLog {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *StudyLogCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *StudyLogCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slc *StudyLogCreate) defaults() {
	if _, ok := slc.mutation.ID(); !ok {
		v := studylog.DefaultID()
		slc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *StudyLogCreate) check() error {
	if _, ok := slc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "StudyLog.start_at"`)}
	}
	if _, ok := slc.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New(`ent: missing required field "StudyLog.end_at"`)}
	}
	if _, ok := slc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "StudyLog.content"`)}
	}
	if _, ok := slc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "StudyLog.user"`)}
	}
	if _, ok := slc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required edge "StudyLog.subject"`)}
	}
	return nil
}

func (slc *StudyLogCreate) sqlSave(ctx context.Context) (*StudyLog, error) {
	if err := slc.check(); err != nil {
		return nil, err
	}
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	slc.mutation.id = &_node.ID
	slc.mutation.done = true
	return _node, nil
}

func (slc *StudyLogCreate) createSpec() (*StudyLog, *sqlgraph.CreateSpec) {
	var (
		_node = &StudyLog{config: slc.config}
		_spec = sqlgraph.NewCreateSpec(studylog.Table, sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID))
	)
	if id, ok := slc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := slc.mutation.StartAt(); ok {
		_spec.SetField(studylog.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := slc.mutation.EndAt(); ok {
		_spec.SetField(studylog.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := slc.mutation.Content(); ok {
		_spec.SetField(studylog.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := slc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_study_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := slc.mutation.SubjectIDs(); len(nodes) > 0 {
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
		_node.subject_study_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := slc.mutation.SharedGroupIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudyLogCreateBulk is the builder for creating many StudyLog entities in bulk.
type StudyLogCreateBulk struct {
	config
	err      error
	builders []*StudyLogCreate
}

// Save creates the StudyLog entities in the database.
func (slcb *StudyLogCreateBulk) Save(ctx context.Context) ([]*StudyLog, error) {
	if slcb.err != nil {
		return nil, slcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*StudyLog, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudyLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *StudyLogCreateBulk) SaveX(ctx context.Context) []*StudyLog {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *StudyLogCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *StudyLogCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}
