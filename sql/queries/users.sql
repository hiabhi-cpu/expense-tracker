-- name: CreateUser :one
INSERT INTO users (
    user_name,
    user_password
) VALUES (
    $1,
    $2
)
RETURNING *;
