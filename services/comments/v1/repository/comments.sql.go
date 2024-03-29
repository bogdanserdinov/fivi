// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: comments.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const create = `-- name: Create :exec
INSERT INTO comments(id, text, post_id, creator_id)
VALUES($1, $2, $3, $4)
`

type CreateParams struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	PostID    uuid.UUID `json:"post_id"`
	CreatorID string    `json:"creator_id"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) error {
	_, err := q.exec(ctx, q.createStmt, create,
		arg.ID,
		arg.Text,
		arg.PostID,
		arg.CreatorID,
	)
	return err
}

const listPostComments = `-- name: ListPostComments :many
SELECT id, text, post_id, creator_id
FROM comments
WHERE post_id = $1
`

func (q *Queries) ListPostComments(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	rows, err := q.query(ctx, q.listPostCommentsStmt, listPostComments, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Text,
			&i.PostID,
			&i.CreatorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
