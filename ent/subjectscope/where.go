// Code generated by ent, DO NOT EDIT.

package subjectscope

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/in-toto/archivista/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldLTE(FieldID, id))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEQ(FieldSubject, v))
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEQ(FieldScope, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldContainsFold(FieldSubject, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldNotIn(FieldScope, vs...))
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldGT(FieldScope, v))
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldGTE(FieldScope, v))
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldLT(FieldScope, v))
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldLTE(FieldScope, v))
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldContains(FieldScope, v))
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldHasPrefix(FieldScope, v))
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldHasSuffix(FieldScope, v))
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldEqualFold(FieldScope, v))
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.SubjectScope {
	return predicate.SubjectScope(sql.FieldContainsFold(FieldScope, v))
}

// HasAttestationPolicy applies the HasEdge predicate on the "attestation_policy" edge.
func HasAttestationPolicy() predicate.SubjectScope {
	return predicate.SubjectScope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AttestationPolicyTable, AttestationPolicyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttestationPolicyWith applies the HasEdge predicate on the "attestation_policy" edge with a given conditions (other predicates).
func HasAttestationPolicyWith(preds ...predicate.AttestationPolicy) predicate.SubjectScope {
	return predicate.SubjectScope(func(s *sql.Selector) {
		step := newAttestationPolicyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SubjectScope) predicate.SubjectScope {
	return predicate.SubjectScope(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SubjectScope) predicate.SubjectScope {
	return predicate.SubjectScope(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SubjectScope) predicate.SubjectScope {
	return predicate.SubjectScope(sql.NotPredicates(p))
}
