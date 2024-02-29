package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SubjectPolicies holds the schema definition for the SubjectPolicies entity.
type SubjectScope struct {
	ent.Schema
}

// Fields of the SubjectScope.
func (SubjectScope) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").NotEmpty(),
		field.String("scope").NotEmpty(),
	}
}

// Edges of the SubjectScope.
func (SubjectScope) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attestation_policy", AttestationPolicy.Type).
			Ref("subject_scopes").
			Unique(),
	}
}

// Edges of the SubjectScope.
func (SubjectScope) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("subject"),
	}
}
