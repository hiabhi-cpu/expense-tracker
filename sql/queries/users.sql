-- name: CreateUser :one
INSERT INTO users (
    user_name,
    user_password
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name like $1;