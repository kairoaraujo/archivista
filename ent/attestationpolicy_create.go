// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/attestationpolicy"
	"github.com/in-toto/archivista/ent/statement"
)

// AttestationPolicyCreate is the builder for creating a AttestationPolicy entity.
type AttestationPolicyCreate struct {
	config
	mutation *AttestationPolicyMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (apc *AttestationPolicyCreate) SetName(s string) *AttestationPolicyCreate {
	apc.mutation.SetName(s)
	return apc
}

// SetID sets the "id" field.
func (apc *AttestationPolicyCreate) SetID(u uuid.UUID) *AttestationPolicyCreate {
	apc.mutation.SetID(u)
	return apc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (apc *AttestationPolicyCreate) SetNillableID(u *uuid.UUID) *AttestationPolicyCreate {
	if u != nil {
		apc.SetID(*u)
	}
	return apc
}

// SetStatementID sets the "statement" edge to the Statement entity by ID.
func (apc *AttestationPolicyCreate) SetStatementID(id uuid.UUID) *AttestationPolicyCreate {
	apc.mutation.SetStatementID(id)
	return apc
}

// SetNillableStatementID sets the "statement" edge to the Statement entity by ID if the given value is not nil.
func (apc *AttestationPolicyCreate) SetNillableStatementID(id *uuid.UUID) *AttestationPolicyCreate {
	if id != nil {
		apc = apc.SetStatementID(*id)
	}
	return apc
}

// SetStatement sets the "statement" edge to the Statement entity.
func (apc *AttestationPolicyCreate) SetStatement(s *Statement) *AttestationPolicyCreate {
	return apc.SetStatementID(s.ID)
}

// Mutation returns the AttestationPolicyMutation object of the builder.
func (apc *AttestationPolicyCreate) Mutation() *AttestationPolicyMutation {
	return apc.mutation
}

// Save creates the AttestationPolicy in the database.
func (apc *AttestationPolicyCreate) Save(ctx context.Context) (*AttestationPolicy, error) {
	apc.defaults()
	return withHooks(ctx, apc.sqlSave, apc.mutation, apc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AttestationPolicyCreate) SaveX(ctx context.Context) *AttestationPolicy {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AttestationPolicyCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AttestationPolicyCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AttestationPolicyCreate) defaults() {
	if _, ok := apc.mutation.ID(); !ok {
		v := attestationpolicy.DefaultID()
		apc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apc *AttestationPolicyCreate) check() error {
	if _, ok := apc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AttestationPolicy.name"`)}
	}
	if v, ok := apc.mutation.Name(); ok {
		if err := attestationpolicy.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AttestationPolicy.name": %w`, err)}
		}
	}
	return nil
}

func (apc *AttestationPolicyCreate) sqlSave(ctx context.Context) (*AttestationPolicy, error) {
	if err := apc.check(); err != nil {
		return nil, err
	}
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
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
	apc.mutation.id = &_node.ID
	apc.mutation.done = true
	return _node, nil
}

func (apc *AttestationPolicyCreate) createSpec() (*AttestationPolicy, *sqlgraph.CreateSpec) {
	var (
		_node = &AttestationPolicy{config: apc.config}
		_spec = sqlgraph.NewCreateSpec(attestationpolicy.Table, sqlgraph.NewFieldSpec(attestationpolicy.FieldID, field.TypeUUID))
	)
	if id, ok := apc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := apc.mutation.Name(); ok {
		_spec.SetField(attestationpolicy.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := apc.mutation.StatementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attestationpolicy.StatementTable,
			Columns: []string{attestationpolicy.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.statement_policy = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttestationPolicyCreateBulk is the builder for creating many AttestationPolicy entities in bulk.
type AttestationPolicyCreateBulk struct {
	config
	err      error
	builders []*AttestationPolicyCreate
}

// Save creates the AttestationPolicy entities in the database.
func (apcb *AttestationPolicyCreateBulk) Save(ctx context.Context) ([]*AttestationPolicy, error) {
	if apcb.err != nil {
		return nil, apcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AttestationPolicy, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttestationPolicyMutation)
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
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AttestationPolicyCreateBulk) SaveX(ctx context.Context) []*AttestationPolicy {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AttestationPolicyCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AttestationPolicyCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}
