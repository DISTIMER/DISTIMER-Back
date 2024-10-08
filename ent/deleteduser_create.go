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
	"pentag.kr/distimer/ent/deleteduser"
)

// DeletedUserCreate is the builder for creating a DeletedUser entity.
type DeletedUserCreate struct {
	config
	mutation *DeletedUserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (duc *DeletedUserCreate) SetName(s string) *DeletedUserCreate {
	duc.mutation.SetName(s)
	return duc
}

// SetOauthID sets the "oauth_id" field.
func (duc *DeletedUserCreate) SetOauthID(s string) *DeletedUserCreate {
	duc.mutation.SetOauthID(s)
	return duc
}

// SetOauthProvider sets the "oauth_provider" field.
func (duc *DeletedUserCreate) SetOauthProvider(i int8) *DeletedUserCreate {
	duc.mutation.SetOauthProvider(i)
	return duc
}

// SetCreatedAt sets the "created_at" field.
func (duc *DeletedUserCreate) SetCreatedAt(t time.Time) *DeletedUserCreate {
	duc.mutation.SetCreatedAt(t)
	return duc
}

// SetDeletedAt sets the "deleted_at" field.
func (duc *DeletedUserCreate) SetDeletedAt(t time.Time) *DeletedUserCreate {
	duc.mutation.SetDeletedAt(t)
	return duc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duc *DeletedUserCreate) SetNillableDeletedAt(t *time.Time) *DeletedUserCreate {
	if t != nil {
		duc.SetDeletedAt(*t)
	}
	return duc
}

// SetID sets the "id" field.
func (duc *DeletedUserCreate) SetID(u uuid.UUID) *DeletedUserCreate {
	duc.mutation.SetID(u)
	return duc
}

// Mutation returns the DeletedUserMutation object of the builder.
func (duc *DeletedUserCreate) Mutation() *DeletedUserMutation {
	return duc.mutation
}

// Save creates the DeletedUser in the database.
func (duc *DeletedUserCreate) Save(ctx context.Context) (*DeletedUser, error) {
	duc.defaults()
	return withHooks(ctx, duc.sqlSave, duc.mutation, duc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (duc *DeletedUserCreate) SaveX(ctx context.Context) *DeletedUser {
	v, err := duc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (duc *DeletedUserCreate) Exec(ctx context.Context) error {
	_, err := duc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duc *DeletedUserCreate) ExecX(ctx context.Context) {
	if err := duc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duc *DeletedUserCreate) defaults() {
	if _, ok := duc.mutation.DeletedAt(); !ok {
		v := deleteduser.DefaultDeletedAt()
		duc.mutation.SetDeletedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duc *DeletedUserCreate) check() error {
	if _, ok := duc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DeletedUser.name"`)}
	}
	if _, ok := duc.mutation.OauthID(); !ok {
		return &ValidationError{Name: "oauth_id", err: errors.New(`ent: missing required field "DeletedUser.oauth_id"`)}
	}
	if _, ok := duc.mutation.OauthProvider(); !ok {
		return &ValidationError{Name: "oauth_provider", err: errors.New(`ent: missing required field "DeletedUser.oauth_provider"`)}
	}
	if _, ok := duc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeletedUser.created_at"`)}
	}
	if _, ok := duc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "DeletedUser.deleted_at"`)}
	}
	return nil
}

func (duc *DeletedUserCreate) sqlSave(ctx context.Context) (*DeletedUser, error) {
	if err := duc.check(); err != nil {
		return nil, err
	}
	_node, _spec := duc.createSpec()
	if err := sqlgraph.CreateNode(ctx, duc.driver, _spec); err != nil {
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
	duc.mutation.id = &_node.ID
	duc.mutation.done = true
	return _node, nil
}

func (duc *DeletedUserCreate) createSpec() (*DeletedUser, *sqlgraph.CreateSpec) {
	var (
		_node = &DeletedUser{config: duc.config}
		_spec = sqlgraph.NewCreateSpec(deleteduser.Table, sqlgraph.NewFieldSpec(deleteduser.FieldID, field.TypeUUID))
	)
	if id, ok := duc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := duc.mutation.Name(); ok {
		_spec.SetField(deleteduser.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := duc.mutation.OauthID(); ok {
		_spec.SetField(deleteduser.FieldOauthID, field.TypeString, value)
		_node.OauthID = value
	}
	if value, ok := duc.mutation.OauthProvider(); ok {
		_spec.SetField(deleteduser.FieldOauthProvider, field.TypeInt8, value)
		_node.OauthProvider = value
	}
	if value, ok := duc.mutation.CreatedAt(); ok {
		_spec.SetField(deleteduser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := duc.mutation.DeletedAt(); ok {
		_spec.SetField(deleteduser.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	return _node, _spec
}

// DeletedUserCreateBulk is the builder for creating many DeletedUser entities in bulk.
type DeletedUserCreateBulk struct {
	config
	err      error
	builders []*DeletedUserCreate
}

// Save creates the DeletedUser entities in the database.
func (ducb *DeletedUserCreateBulk) Save(ctx context.Context) ([]*DeletedUser, error) {
	if ducb.err != nil {
		return nil, ducb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ducb.builders))
	nodes := make([]*DeletedUser, len(ducb.builders))
	mutators := make([]Mutator, len(ducb.builders))
	for i := range ducb.builders {
		func(i int, root context.Context) {
			builder := ducb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeletedUserMutation)
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
					_, err = mutators[i+1].Mutate(root, ducb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ducb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ducb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ducb *DeletedUserCreateBulk) SaveX(ctx context.Context) []*DeletedUser {
	v, err := ducb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ducb *DeletedUserCreateBulk) Exec(ctx context.Context) error {
	_, err := ducb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ducb *DeletedUserCreateBulk) ExecX(ctx context.Context) {
	if err := ducb.Exec(ctx); err != nil {
		panic(err)
	}
}
