// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/in-toto/archivista/ent/attestation"
	"github.com/in-toto/archivista/ent/attestationcollection"
	"github.com/in-toto/archivista/ent/attestationpolicy"
	"github.com/in-toto/archivista/ent/dsse"
	"github.com/in-toto/archivista/ent/payloaddigest"
	"github.com/in-toto/archivista/ent/signature"
	"github.com/in-toto/archivista/ent/statement"
	"github.com/in-toto/archivista/ent/subject"
	"github.com/in-toto/archivista/ent/subjectdigest"
	"github.com/in-toto/archivista/ent/subjectscope"
	"github.com/in-toto/archivista/ent/timestamp"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *Attestation) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *AttestationCollection) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *AttestationPolicy) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Dsse) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *PayloadDigest) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Signature) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Statement) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Subject) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *SubjectDigest) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *SubjectScope) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Timestamp) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case attestation.Table:
		query := c.Attestation.Query().
			Where(attestation.ID(id))
		query, err := query.CollectFields(ctx, "Attestation")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case attestationcollection.Table:
		query := c.AttestationCollection.Query().
			Where(attestationcollection.ID(id))
		query, err := query.CollectFields(ctx, "AttestationCollection")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case attestationpolicy.Table:
		query := c.AttestationPolicy.Query().
			Where(attestationpolicy.ID(id))
		query, err := query.CollectFields(ctx, "AttestationPolicy")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case dsse.Table:
		query := c.Dsse.Query().
			Where(dsse.ID(id))
		query, err := query.CollectFields(ctx, "Dsse")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case payloaddigest.Table:
		query := c.PayloadDigest.Query().
			Where(payloaddigest.ID(id))
		query, err := query.CollectFields(ctx, "PayloadDigest")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case signature.Table:
		query := c.Signature.Query().
			Where(signature.ID(id))
		query, err := query.CollectFields(ctx, "Signature")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case statement.Table:
		query := c.Statement.Query().
			Where(statement.ID(id))
		query, err := query.CollectFields(ctx, "Statement")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subject.Table:
		query := c.Subject.Query().
			Where(subject.ID(id))
		query, err := query.CollectFields(ctx, "Subject")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subjectdigest.Table:
		query := c.SubjectDigest.Query().
			Where(subjectdigest.ID(id))
		query, err := query.CollectFields(ctx, "SubjectDigest")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subjectscope.Table:
		query := c.SubjectScope.Query().
			Where(subjectscope.ID(id))
		query, err := query.CollectFields(ctx, "SubjectScope")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case timestamp.Table:
		query := c.Timestamp.Query().
			Where(timestamp.ID(id))
		query, err := query.CollectFields(ctx, "Timestamp")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case attestation.Table:
		query := c.Attestation.Query().
			Where(attestation.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Attestation")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case attestationcollection.Table:
		query := c.AttestationCollection.Query().
			Where(attestationcollection.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AttestationCollection")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case attestationpolicy.Table:
		query := c.AttestationPolicy.Query().
			Where(attestationpolicy.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AttestationPolicy")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case dsse.Table:
		query := c.Dsse.Query().
			Where(dsse.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Dsse")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case payloaddigest.Table:
		query := c.PayloadDigest.Query().
			Where(payloaddigest.IDIn(ids...))
		query, err := query.CollectFields(ctx, "PayloadDigest")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case signature.Table:
		query := c.Signature.Query().
			Where(signature.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Signature")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case statement.Table:
		query := c.Statement.Query().
			Where(statement.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Statement")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subject.Table:
		query := c.Subject.Query().
			Where(subject.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Subject")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subjectdigest.Table:
		query := c.SubjectDigest.Query().
			Where(subjectdigest.IDIn(ids...))
		query, err := query.CollectFields(ctx, "SubjectDigest")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subjectscope.Table:
		query := c.SubjectScope.Query().
			Where(subjectscope.IDIn(ids...))
		query, err := query.CollectFields(ctx, "SubjectScope")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case timestamp.Table:
		query := c.Timestamp.Query().
			Where(timestamp.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Timestamp")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
