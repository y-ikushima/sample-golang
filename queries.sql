-- name: GetUser :one
select u.id,round(( random() * (1 - 10000) )::numeric, 0) + 10000 as num from users u where u.id = $1;