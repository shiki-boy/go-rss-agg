// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Todo struct {
	ID        int32
	Title     pgtype.Text
	Done      pgtype.Bool
	CreatedAt pgtype.Timestamp
}