package tree

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dgraph-io/ristretto"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PgRistrettoStore uses PostgreSQL with a ristretto cache in front.
type PgRistrettoStore struct {
	db             *pgxpool.Pool
	dbTx           pgx.Tx
	tableName      string
	constraintName string
	cache          *ristretto.Cache
}

// NewPgRistrettoStore creates an instance of PgRistrettoStore.
func NewPgRistrettoStore(db *pgxpool.Pool, cache *ristretto.Cache) *PgRistrettoStore {
	return &PgRistrettoStore{
		db:             db,
		tableName:      merkleTreeTable,
		constraintName: mtConstraint,
		cache:          cache,
	}
}

// NewPgRistrettoSCCodeStore creates an instance of PgRistrettoStore.
func NewPgRistrettoSCCodeStore(db *pgxpool.Pool, cache *ristretto.Cache) *PgRistrettoStore {
	return &PgRistrettoStore{
		db:             db,
		tableName:      scCodeTreeTable,
		constraintName: scCodeConstraint,
		cache:          cache,
	}
}

// BeginDBTransaction starts a transaction block
func (p *PgRistrettoStore) BeginDBTransaction(ctx context.Context) error {
	dbTx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	p.dbTx = dbTx
	return nil
}

// Commit commits a db transaction
func (p *PgRistrettoStore) Commit(ctx context.Context) error {
	if p.dbTx != nil {
		err := p.dbTx.Commit(ctx)
		p.dbTx = nil
		return err
	}

	return ErrNilDBTransaction
}

// Rollback rollbacks a db transaction
func (p *PgRistrettoStore) Rollback(ctx context.Context) error {
	if p.dbTx != nil {
		err := p.dbTx.Rollback(ctx)
		p.dbTx = nil
		return err
	}

	return ErrNilDBTransaction
}

func (p *PgRistrettoStore) exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
	if p.dbTx != nil {
		return p.dbTx.Exec(ctx, sql, arguments...)
	}
	return p.db.Exec(ctx, sql, arguments...)
}

func (p *PgRistrettoStore) queryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if p.dbTx != nil {
		return p.dbTx.QueryRow(ctx, sql, args...)
	}
	return p.db.QueryRow(ctx, sql, args...)
}

// Get gets value of key, first trying the cache, then the db.
func (p *PgRistrettoStore) Get(ctx context.Context, key []byte) ([]byte, error) {
	value, found := p.cache.Get(key)
	if found {
		data, ok := value.([]byte)
		if !ok {
			return nil, fmt.Errorf("Could not cast data to []byte for key %q", key)
		}
		return data, nil
	}
	var data []byte
	err := p.queryRow(ctx, fmt.Sprintf(getNodeByKeySQL, p.tableName), key).Scan(&data)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	p.cache.Set(key, data, cacheDefaultCost)
	return data, nil
}

// Set inserts a key-value pair into the db.
// If record with such a key already exists its assumed that the value is correct,
// because it's a reverse hash table, and the key is a hash of the value.
func (p *PgRistrettoStore) Set(ctx context.Context, key []byte, value []byte) error {
	_, err := p.exec(ctx, fmt.Sprintf(setNodeByKeySQL, p.tableName, p.constraintName), key, value)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil
		}
		return err
	}
	return nil
}

// Reset clears the db and the cache.
func (p *PgRistrettoStore) Reset() error {
	p.cache.Clear()

	_, err := p.exec(context.Background(), fmt.Sprintf("TRUNCATE TABLE %s;", p.tableName))
	return err
}