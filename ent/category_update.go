// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/predicate"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/user"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableName(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetOrder sets the "order" field.
func (cu *CategoryUpdate) SetOrder(i int8) *CategoryUpdate {
	cu.mutation.ResetOrder()
	cu.mutation.SetOrder(i)
	return cu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableOrder(i *int8) *CategoryUpdate {
	if i != nil {
		cu.SetOrder(*i)
	}
	return cu
}

// AddOrder adds i to the "order" field.
func (cu *CategoryUpdate) AddOrder(i int8) *CategoryUpdate {
	cu.mutation.AddOrder(i)
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CategoryUpdate) SetUserID(id uuid.UUID) *CategoryUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CategoryUpdate) SetUser(u *User) *CategoryUpdate {
	return cu.SetUserID(u.ID)
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (cu *CategoryUpdate) AddSubjectIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.AddSubjectIDs(ids...)
	return cu
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (cu *CategoryUpdate) AddSubjects(s ...*Subject) *CategoryUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSubjectIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CategoryUpdate) ClearUser() *CategoryUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (cu *CategoryUpdate) ClearSubjects() *CategoryUpdate {
	cu.mutation.ClearSubjects()
	return cu
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (cu *CategoryUpdate) RemoveSubjectIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.RemoveSubjectIDs(ids...)
	return cu
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (cu *CategoryUpdate) RemoveSubjects(s ...*Subject) *CategoryUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSubjectIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CategoryUpdate) check() error {
	if _, ok := cu.mutation.UserID(); cu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Category.user"`)
	}
	return nil
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Order(); ok {
		_spec.SetField(category.FieldOrder, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedOrder(); ok {
		_spec.AddField(category.FieldOrder, field.TypeInt8, value)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.UserTable,
			Columns: []string{category.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.UserTable,
			Columns: []string{category.UserColumn},
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
	if cu.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SubjectsTable,
			Columns: []string{category.SubjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !cu.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SubjectsTable,
			Columns: []string{category.SubjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SubjectsTable,
			Columns: []string{category.SubjectsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableName(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetOrder sets the "order" field.
func (cuo *CategoryUpdateOne) SetOrder(i int8) *CategoryUpdateOne {
	cuo.mutation.ResetOrder()
	cuo.mutation.SetOrder(i)
	return cuo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableOrder(i *int8) *CategoryUpdateOne {
	if i != nil {
		cuo.SetOrder(*i)
	}
	return cuo
}

// AddOrder adds i to the "order" field.
func (cuo *CategoryUpdateOne) AddOrder(i int8) *CategoryUpdateOne {
	cuo.mutation.AddOrder(i)
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CategoryUpdateOne) SetUserID(id uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CategoryUpdateOne) SetUser(u *User) *CategoryUpdateOne {
	return cuo.SetUserID(u.ID)
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (cuo *CategoryUpdateOne) AddSubjectIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.AddSubjectIDs(ids...)
	return cuo
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (cuo *CategoryUpdateOne) AddSubjects(s ...*Subject) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSubjectIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CategoryUpdateOne) ClearUser() *CategoryUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (cuo *CategoryUpdateOne) ClearSubjects() *CategoryUpdateOne {
	cuo.mutation.ClearSubjects()
	return cuo
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (cuo *CategoryUpdateOne) RemoveSubjectIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.RemoveSubjectIDs(ids...)
	return cuo
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (cuo *CategoryUpdateOne) RemoveSubjects(s ...*Subject) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSubjectIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CategoryUpdateOne) check() error {
	if _, ok := cuo.mutation.UserID(); cuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Category.user"`)
	}
	return nil
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Order(); ok {
		_spec.SetField(category.FieldOrder, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedOrder(); ok {
		_spec.AddField(category.FieldOrder, field.TypeInt8, value)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.UserTable,
			Columns: []string{category.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.UserTable,
			Columns: []string{category.UserColumn},
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
	if cuo.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SubjectsTable,
			Columns: []string{category.SubjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !cuo.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SubjectsTable,
			Columns: []string{category.SubjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SubjectsTable,
			Columns: []string{category.SubjectsColumn},
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
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
