-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $1,
    dob  = $2
WHERE id = $3
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;