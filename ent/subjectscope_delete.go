// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/in-toto/archivista/ent/predicate"
	"github.com/in-toto/archivista/ent/subjectscope"
)

// SubjectScopeDelete is the builder for deleting a SubjectScope entity.
type SubjectScopeDelete struct {
	config
	hooks    []Hook
	mutation *SubjectScopeMutation
}

// Where appends a list predicates to the SubjectScopeDelete builder.
func (ssd *SubjectScopeDelete) Where(ps ...predicate.SubjectScope) *SubjectScopeDelete {
	ssd.mutation.Where(ps...)
	return ssd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ssd *SubjectScopeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ssd.sqlExec, ssd.mutation, ssd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ssd *SubjectScopeDelete) ExecX(ctx context.Context) int {
	n, err := ssd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ssd *SubjectScopeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subjectscope.Table, sqlgraph.NewFieldSpec(subjectscope.FieldID, field.TypeInt))
	if ps := ssd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ssd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ssd.mutation.done = true
	return affected, err
}

// SubjectScopeDeleteOne is the builder for deleting a single SubjectScope entity.
type SubjectScopeDeleteOne struct {
	ssd *SubjectScopeDelete
}

// Where appends a list predicates to the SubjectScopeDelete builder.
func (ssdo *SubjectScopeDeleteOne) Where(ps ...predicate.SubjectScope) *SubjectScopeDeleteOne {
	ssdo.ssd.mutation.Where(ps...)
	return ssdo
}

// Exec executes the deletion query.
func (ssdo *SubjectScopeDeleteOne) Exec(ctx context.Context) error {
	n, err := ssdo.ssd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subjectscope.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ssdo *SubjectScopeDeleteOne) ExecX(ctx context.Context) {
	if err := ssdo.Exec(ctx); err != nil {
		panic(err)
	}
}
