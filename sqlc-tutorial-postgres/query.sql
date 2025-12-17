-- name: GetUsers :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUsers :one
INSERT INTO users (
  name, dob
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUsers :one
UPDATE users
  set name = $2,
  dob = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE id = $1;