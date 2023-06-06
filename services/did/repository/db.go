// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createDIDMappingStmt, err = db.PrepareContext(ctx, createDIDMapping); err != nil {
		return nil, fmt.Errorf("error preparing query CreateDIDMapping: %w", err)
	}
	if q.deleteDIDMappingStmt, err = db.PrepareContext(ctx, deleteDIDMapping); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteDIDMapping: %w", err)
	}
	if q.getDIDMappingStmt, err = db.PrepareContext(ctx, getDIDMapping); err != nil {
		return nil, fmt.Errorf("error preparing query GetDIDMapping: %w", err)
	}
	if q.getDIDMessageStmt, err = db.PrepareContext(ctx, getDIDMessage); err != nil {
		return nil, fmt.Errorf("error preparing query GetDIDMessage: %w", err)
	}
	if q.upsertDIDMessageStmt, err = db.PrepareContext(ctx, upsertDIDMessage); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertDIDMessage: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createDIDMappingStmt != nil {
		if cerr := q.createDIDMappingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createDIDMappingStmt: %w", cerr)
		}
	}
	if q.deleteDIDMappingStmt != nil {
		if cerr := q.deleteDIDMappingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteDIDMappingStmt: %w", cerr)
		}
	}
	if q.getDIDMappingStmt != nil {
		if cerr := q.getDIDMappingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDIDMappingStmt: %w", cerr)
		}
	}
	if q.getDIDMessageStmt != nil {
		if cerr := q.getDIDMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDIDMessageStmt: %w", cerr)
		}
	}
	if q.upsertDIDMessageStmt != nil {
		if cerr := q.upsertDIDMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertDIDMessageStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	createDIDMappingStmt *sql.Stmt
	deleteDIDMappingStmt *sql.Stmt
	getDIDMappingStmt    *sql.Stmt
	getDIDMessageStmt    *sql.Stmt
	upsertDIDMessageStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createDIDMappingStmt: q.createDIDMappingStmt,
		deleteDIDMappingStmt: q.deleteDIDMappingStmt,
		getDIDMappingStmt:    q.getDIDMappingStmt,
		getDIDMessageStmt:    q.getDIDMessageStmt,
		upsertDIDMessageStmt: q.upsertDIDMessageStmt,
	}
}