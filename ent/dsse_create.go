// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/dsse"
	"github.com/in-toto/archivista/ent/payloaddigest"
	"github.com/in-toto/archivista/ent/signature"
	"github.com/in-toto/archivista/ent/statement"
)

// DsseCreate is the builder for creating a Dsse entity.
type DsseCreate struct {
	config
	mutation *DsseMutation
	hooks    []Hook
}

// SetGitoidSha256 sets the "gitoid_sha256" field.
func (dc *DsseCreate) SetGitoidSha256(s string) *DsseCreate {
	dc.mutation.SetGitoidSha256(s)
	return dc
}

// SetPayloadType sets the "payload_type" field.
func (dc *DsseCreate) SetPayloadType(s string) *DsseCreate {
	dc.mutation.SetPayloadType(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DsseCreate) SetID(u uuid.UUID) *DsseCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DsseCreate) SetNillableID(u *uuid.UUID) *DsseCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// SetStatementID sets the "statement" edge to the Statement entity by ID.
func (dc *DsseCreate) SetStatementID(id uuid.UUID) *DsseCreate {
	dc.mutation.SetStatementID(id)
	return dc
}

// SetNillableStatementID sets the "statement" edge to the Statement entity by ID if the given value is not nil.
func (dc *DsseCreate) SetNillableStatementID(id *uuid.UUID) *DsseCreate {
	if id != nil {
		dc = dc.SetStatementID(*id)
	}
	return dc
}

// SetStatement sets the "statement" edge to the Statement entity.
func (dc *DsseCreate) SetStatement(s *Statement) *DsseCreate {
	return dc.SetStatementID(s.ID)
}

// AddSignatureIDs adds the "signatures" edge to the Signature entity by IDs.
func (dc *DsseCreate) AddSignatureIDs(ids ...uuid.UUID) *DsseCreate {
	dc.mutation.AddSignatureIDs(ids...)
	return dc
}

// AddSignatures adds the "signatures" edges to the Signature entity.
func (dc *DsseCreate) AddSignatures(s ...*Signature) *DsseCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddSignatureIDs(ids...)
}

// AddPayloadDigestIDs adds the "payload_digests" edge to the PayloadDigest entity by IDs.
func (dc *DsseCreate) AddPayloadDigestIDs(ids ...uuid.UUID) *DsseCreate {
	dc.mutation.AddPayloadDigestIDs(ids...)
	return dc
}

// AddPayloadDigests adds the "payload_digests" edges to the PayloadDigest entity.
func (dc *DsseCreate) AddPayloadDigests(p ...*PayloadDigest) *DsseCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddPayloadDigestIDs(ids...)
}

// Mutation returns the DsseMutation object of the builder.
func (dc *DsseCreate) Mutation() *DsseMutation {
	return dc.mutation
}

// Save creates the Dsse in the database.
func (dc *DsseCreate) Save(ctx context.Context) (*Dsse, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DsseCreate) SaveX(ctx context.Context) *Dsse {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DsseCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DsseCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DsseCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := dsse.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DsseCreate) check() error {
	if _, ok := dc.mutation.GitoidSha256(); !ok {
		return &ValidationError{Name: "gitoid_sha256", err: errors.New(`ent: missing required field "Dsse.gitoid_sha256"`)}
	}
	if v, ok := dc.mutation.GitoidSha256(); ok {
		if err := dsse.GitoidSha256Validator(v); err != nil {
			return &ValidationError{Name: "gitoid_sha256", err: fmt.Errorf(`ent: validator failed for field "Dsse.gitoid_sha256": %w`, err)}
		}
	}
	if _, ok := dc.mutation.PayloadType(); !ok {
		return &ValidationError{Name: "payload_type", err: errors.New(`ent: missing required field "Dsse.payload_type"`)}
	}
	if v, ok := dc.mutation.PayloadType(); ok {
		if err := dsse.PayloadTypeValidator(v); err != nil {
			return &ValidationError{Name: "payload_type", err: fmt.Errorf(`ent: validator failed for field "Dsse.payload_type": %w`, err)}
		}
	}
	return nil
}

func (dc *DsseCreate) sqlSave(ctx context.Context) (*Dsse, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DsseCreate) createSpec() (*Dsse, *sqlgraph.CreateSpec) {
	var (
		_node = &Dsse{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dsse.Table, sqlgraph.NewFieldSpec(dsse.FieldID, field.TypeUUID))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.GitoidSha256(); ok {
		_spec.SetField(dsse.FieldGitoidSha256, field.TypeString, value)
		_node.GitoidSha256 = value
	}
	if value, ok := dc.mutation.PayloadType(); ok {
		_spec.SetField(dsse.FieldPayloadType, field.TypeString, value)
		_node.PayloadType = value
	}
	if nodes := dc.mutation.StatementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dsse.StatementTable,
			Columns: []string{dsse.StatementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.dsse_statement = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.SignaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dsse.SignaturesTable,
			Columns: []string{dsse.SignaturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signature.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.PayloadDigestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dsse.PayloadDigestsTable,
			Columns: []string{dsse.PayloadDigestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payloaddigest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DsseCreateBulk is the builder for creating many Dsse entities in bulk.
type DsseCreateBulk struct {
	config
	err      error
	builders []*DsseCreate
}

// Save creates the Dsse entities in the database.
func (dcb *DsseCreateBulk) Save(ctx context.Context) ([]*Dsse, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dsse, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DsseMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DsseCreateBulk) SaveX(ctx context.Context) []*Dsse {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DsseCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DsseCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
