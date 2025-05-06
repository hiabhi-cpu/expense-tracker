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

-- name: GetMoneyPerId :one
SELECT * FROM money
where mon_id =$1;

-- name: UpdateMoney :exec
UPDATE money 
set amt = $3
where mon_id =$1 and user_id=$2;

-- name: DeleteMoney :exec
DELETE FROM money
where mon_id = $1 and user_id=$2;