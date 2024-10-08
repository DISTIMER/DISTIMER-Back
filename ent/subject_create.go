// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
)

// SubjectCreate is the builder for creating a Subject entity.
type SubjectCreate struct {
	config
	mutation *SubjectMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SubjectCreate) SetName(s string) *SubjectCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetColor sets the "color" field.
func (sc *SubjectCreate) SetColor(s string) *SubjectCreate {
	sc.mutation.SetColor(s)
	return sc
}

// SetOrder sets the "order" field.
func (sc *SubjectCreate) SetOrder(i int8) *SubjectCreate {
	sc.mutation.SetOrder(i)
	return sc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableOrder(i *int8) *SubjectCreate {
	if i != nil {
		sc.SetOrder(*i)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SubjectCreate) SetID(u uuid.UUID) *SubjectCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableID(u *uuid.UUID) *SubjectCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (sc *SubjectCreate) SetCategoryID(id uuid.UUID) *SubjectCreate {
	sc.mutation.SetCategoryID(id)
	return sc
}

// SetCategory sets the "category" edge to the Category entity.
func (sc *SubjectCreate) SetCategory(c *Category) *SubjectCreate {
	return sc.SetCategoryID(c.ID)
}

// AddStudyLogIDs adds the "study_logs" edge to the StudyLog entity by IDs.
func (sc *SubjectCreate) AddStudyLogIDs(ids ...uuid.UUID) *SubjectCreate {
	sc.mutation.AddStudyLogIDs(ids...)
	return sc
}

// AddStudyLogs adds the "study_logs" edges to the StudyLog entity.
func (sc *SubjectCreate) AddStudyLogs(s ...*StudyLog) *SubjectCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStudyLogIDs(ids...)
}

// AddTimerIDs adds the "timers" edge to the Timer entity by IDs.
func (sc *SubjectCreate) AddTimerIDs(ids ...uuid.UUID) *SubjectCreate {
	sc.mutation.AddTimerIDs(ids...)
	return sc
}

// AddTimers adds the "timers" edges to the Timer entity.
func (sc *SubjectCreate) AddTimers(t ...*Timer) *SubjectCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTimerIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (sc *SubjectCreate) Mutation() *SubjectMutation {
	return sc.mutation
}

// Save creates the Subject in the database.
func (sc *SubjectCreate) Save(ctx context.Context) (*Subject, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubjectCreate) SaveX(ctx context.Context) *Subject {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubjectCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubjectCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubjectCreate) defaults() {
	if _, ok := sc.mutation.Order(); !ok {
		v := subject.DefaultOrder
		sc.mutation.SetOrder(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := subject.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubjectCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Subject.name"`)}
	}
	if _, ok := sc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "Subject.color"`)}
	}
	if _, ok := sc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Subject.order"`)}
	}
	if _, ok := sc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required edge "Subject.category"`)}
	}
	return nil
}

func (sc *SubjectCreate) sqlSave(ctx context.Context) (*Subject, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubjectCreate) createSpec() (*Subject, *sqlgraph.CreateSpec) {
	var (
		_node = &Subject{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subject.Table, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Color(); ok {
		_spec.SetField(subject.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if value, ok := sc.mutation.Order(); ok {
		_spec.SetField(subject.FieldOrder, field.TypeInt8, value)
		_node.Order = value
	}
	if nodes := sc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.CategoryTable,
			Columns: []string{subject.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_subjects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudyLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TimersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubjectCreateBulk is the builder for creating many Subject entities in bulk.
type SubjectCreateBulk struct {
	config
	err      error
	builders []*SubjectCreate
}

// Save creates the Subject entities in the database.
func (scb *SubjectCreateBulk) Save(ctx context.Context) ([]*Subject, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subject, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubjectCreateBulk) SaveX(ctx context.Context) []*Subject {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubjectCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubjectCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
