// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package repository

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Payload   string    `json:"payload"`
	CreatorID uuid.UUID `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
}