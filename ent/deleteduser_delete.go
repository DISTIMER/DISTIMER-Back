// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"pentag.kr/distimer/ent/deleteduser"
	"pentag.kr/distimer/ent/predicate"
)

// DeletedUserDelete is the builder for deleting a DeletedUser entity.
type DeletedUserDelete struct {
	config
	hooks    []Hook
	mutation *DeletedUserMutation
}

// Where appends a list predicates to the DeletedUserDelete builder.
func (dud *DeletedUserDelete) Where(ps ...predicate.DeletedUser) *DeletedUserDelete {
	dud.mutation.Where(ps...)
	return dud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dud *DeletedUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dud.sqlExec, dud.mutation, dud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dud *DeletedUserDelete) ExecX(ctx context.Context) int {
	n, err := dud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dud *DeletedUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(deleteduser.Table, sqlgraph.NewFieldSpec(deleteduser.FieldID, field.TypeUUID))
	if ps := dud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dud.mutation.done = true
	return affected, err
}

// DeletedUserDeleteOne is the builder for deleting a single DeletedUser entity.
type DeletedUserDeleteOne struct {
	dud *DeletedUserDelete
}

// Where appends a list predicates to the DeletedUserDelete builder.
func (dudo *DeletedUserDeleteOne) Where(ps ...predicate.DeletedUser) *DeletedUserDeleteOne {
	dudo.dud.mutation.Where(ps...)
	return dudo
}

// Exec executes the deletion query.
func (dudo *DeletedUserDeleteOne) Exec(ctx context.Context) error {
	n, err := dudo.dud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deleteduser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dudo *DeletedUserDeleteOne) ExecX(ctx context.Context) {
	if err := dudo.Exec(ctx); err != nil {
		panic(err)
	}
}
