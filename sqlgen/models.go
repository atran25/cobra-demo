// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlgen

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Character struct {
	CharacterID   string
	ServerID      string
	CharacterName string
	JobID         string
	JobGrowID     string
	Rank          pgtype.Int4
}