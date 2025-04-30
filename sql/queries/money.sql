-- name: CreateMoney :one
INSERT INTO money (
    mon_desc,
    amt,
    user_id ,
    mon_date
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetAllMoney :many
SELECT * FROM money
WHERE user_id = $1 
ORDER BY mon_date;