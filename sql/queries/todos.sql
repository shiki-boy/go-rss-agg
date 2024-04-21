-- name: CreateTodo :one
INSERT INTO Todo (title, done) 
VALUES ($1, $2)
RETURNING *;