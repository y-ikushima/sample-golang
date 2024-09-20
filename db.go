// db.go
package main

import (
	"context"
	"sample-golang/db/sqlc"
)


func getUsers(q *sqlc.Queries, id int32) (*sqlc.GetUserRow, error) {
	user, err := q.GetUser(context.Background(),id) // 生成されたクエリを呼び出す
	if err != nil {
		return nil, err
	}
	return &user, nil
}