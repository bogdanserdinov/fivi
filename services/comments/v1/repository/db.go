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
	if q.createStmt, err = db.PrepareContext(ctx, create); err != nil {
		return nil, fmt.Errorf("error preparing query Create: %w", err)
	}
	if q.listPostCommentsStmt, err = db.PrepareContext(ctx, listPostComments); err != nil {
		return nil, fmt.Errorf("error preparing query ListPostComments: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createStmt != nil {
		if cerr := q.createStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStmt: %w", cerr)
		}
	}
	if q.listPostCommentsStmt != nil {
		if cerr := q.listPostCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPostCommentsStmt: %w", cerr)
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
	db DBTX
	tx *sql.Tx
	createStmt           *sql.Stmt
	listPostCommentsStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createStmt:           q.createStmt,
		listPostCommentsStmt: q.listPostCommentsStmt,
	}
}
