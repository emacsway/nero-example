// Code generated by nero, DO NOT EDIT.
package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sf9v/nero"
	"github.com/sf9v/nero-example/model"
)

// Repository is a repository for Product
type Repository interface {
	// Tx begins a new transaction
	Tx(context.Context) (nero.Tx, error)
	// Create creates a new Product
	Create(context.Context, *Creator) (id int64, err error)
	// CreateTx creates a new type .Type.Name}} inside a transaction
	CreateTx(context.Context, nero.Tx, *Creator) (id int64, err error)
	// CreateMany creates many Product
	CreateMany(context.Context, ...*Creator) error
	// CreateManyTx creates many Product inside a transaction
	CreateManyTx(context.Context, nero.Tx, ...*Creator) error
	// Query queries many Product
	Query(context.Context, *Queryer) ([]*model.Product, error)
	// QueryTx queries many {0  <nil> <nil>} inside a transaction
	QueryTx(context.Context, nero.Tx, *Queryer) ([]*model.Product, error)
	// QueryOne queries one Product
	QueryOne(context.Context, *Queryer) (*model.Product, error)
	// QueryOneTx queries one Product inside a transaction
	QueryOneTx(context.Context, nero.Tx, *Queryer) (*model.Product, error)
	// Update updates Product
	Update(context.Context, *Updater) (rowsAffected int64, err error)
	// UpdateTx updates Product inside a transaction
	UpdateTx(context.Context, nero.Tx, *Updater) (rowsAffected int64, err error)
	// Delete deletes Product
	Delete(context.Context, *Deleter) (rowsAffected int64, err error)
	// Delete deletes Product inside a transaction
	DeleteTx(context.Context, nero.Tx, *Deleter) (rowsAffected int64, err error)
	// Aggregate performs aggregate operations
	Aggregate(context.Context, *Aggregator) error
	// Aggregate performs aggregate operations inside a transaction
	AggregateTx(context.Context, nero.Tx, *Aggregator) error
}

// Creator is a create builder for Product
type Creator struct {
	name      string
	updatedAt *time.Time
}

// NewCreator is a factory for Creator
func NewCreator() *Creator {
	return &Creator{}
}

// Name is a setter for name
func (c *Creator) Name(name string) *Creator {
	c.name = name
	return c
}

// UpdatedAt is a setter for updatedAt
func (c *Creator) UpdatedAt(updatedAt *time.Time) *Creator {
	c.updatedAt = updatedAt
	return c
}

// Queryer is a query builder for Product
type Queryer struct {
	limit  uint
	offset uint
	pfs    []PredFunc
	sfs    []SortFunc
}

// NewQueryer is a factory for Queryer
func NewQueryer() *Queryer {
	return &Queryer{}
}

// Where adds predicates to the query
func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

// Sort adds sorting to the query
func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

// Limit adds limit to the query
func (q *Queryer) Limit(limit uint) *Queryer {
	q.limit = limit
	return q
}

// Offset adds offset to the query
func (q *Queryer) Offset(offset uint) *Queryer {
	q.offset = offset
	return q
}

// Updater is an update builder for Product
type Updater struct {
	name      string
	updatedAt *time.Time
	pfs       []PredFunc
}

// NewUpdater is a factory for Updater
func NewUpdater() *Updater {
	return &Updater{}
}

// Name is a setter for name
func (c *Updater) Name(name string) *Updater {
	c.name = name
	return c
}

// UpdatedAt is a setter for updatedAt
func (c *Updater) UpdatedAt(updatedAt *time.Time) *Updater {
	c.updatedAt = updatedAt
	return c
}

// Where adds predicates to the update
func (u *Updater) Where(pfs ...PredFunc) *Updater {
	u.pfs = append(u.pfs, pfs...)
	return u
}

// Deleter is an update builder for Product
type Deleter struct {
	pfs []PredFunc
}

// NewDeleter is a factory for Deleter
func NewDeleter() *Deleter {
	return &Deleter{}
}

// Where adds predicates to the delete
func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}

// Aggregator is an aggregate builder for Product
type Aggregator struct {
	v      interface{}
	aggfs  []AggFunc
	pfs    []PredFunc
	sfs    []SortFunc
	groups []Column
}

// NewAggregator is a factory for Aggregator
// v must be an array of struct
func NewAggregator(v interface{}) *Aggregator {
	return &Aggregator{
		v: v,
	}
}

// Aggregate adds aggregate functions to the aggregate
func (a *Aggregator) Aggregate(aggfs ...AggFunc) *Aggregator {
	a.aggfs = append(a.aggfs, aggfs...)
	return a
}

// Where adds predicates to the aggregate
func (a *Aggregator) Where(pfs ...PredFunc) *Aggregator {
	a.pfs = append(a.pfs, pfs...)
	return a
}

// Sort adds sorting to the aggregate
func (a *Aggregator) Sort(sfs ...SortFunc) *Aggregator {
	a.sfs = append(a.sfs, sfs...)
	return a
}

// Group adds grouping to the aggregate
func (a *Aggregator) Group(cols ...Column) *Aggregator {
	a.groups = append(a.groups, cols...)
	return a
}

func rollback(tx nero.Tx, err error) error {
	rerr := tx.Rollback()
	if rerr != nil {
		err = errors.Wrapf(err, "rollback error: %v", rerr)
	}
	return err
}
