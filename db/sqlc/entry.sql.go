// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
  accounts_id,
  amount
) VALUES (
  $1, $2
)
RETURNING id, accounts_id, amount, created_at
`

type CreateEntryParams struct {
	AccountsID int64 `json:"accounts_id"`
	Amount     int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (*Entries, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountsID, arg.Amount)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountsID,
		&i.Amount,
		&i.CreatedAt,
	)
	return &i, err
}

const getEntry = `-- name: GetEntry :one
SELECT id, accounts_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (*Entries, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountsID,
		&i.Amount,
		&i.CreatedAt,
	)
	return &i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, accounts_id, amount, created_at FROM entries
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]*Entries, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Entries
	for rows.Next() {
		var i Entries
		if err := rows.Scan(
			&i.ID,
			&i.AccountsID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
