// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/predicate"
)

// AffiliationDelete is the builder for deleting a Affiliation entity.
type AffiliationDelete struct {
	config
	hooks    []Hook
	mutation *AffiliationMutation
}

// Where appends a list predicates to the AffiliationDelete builder.
func (ad *AffiliationDelete) Where(ps ...predicate.Affiliation) *AffiliationDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AffiliationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AffiliationDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AffiliationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(affiliation.Table, nil)
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AffiliationDeleteOne is the builder for deleting a single Affiliation entity.
type AffiliationDeleteOne struct {
	ad *AffiliationDelete
}

// Where appends a list predicates to the AffiliationDelete builder.
func (ado *AffiliationDeleteOne) Where(ps ...predicate.Affiliation) *AffiliationDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AffiliationDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{affiliation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AffiliationDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
