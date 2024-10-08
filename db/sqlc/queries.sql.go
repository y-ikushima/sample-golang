// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package sqlc

import (
	"context"
)

const getUser = `-- name: GetUser :one
select u.id,round(( random() * (1 - 10000) )::numeric, 0) + 10000 as num from users u where u.id = $1
`

type GetUserRow struct {
	ID  int32
	Num int32
}

func (q *Queries) GetUser(ctx context.Context, id int32) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(&i.ID, &i.Num)
	return i, err
}
