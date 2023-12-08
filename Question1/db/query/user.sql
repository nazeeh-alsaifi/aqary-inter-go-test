-- name: CreateUser :one
INSERT INTO users (
  name,
  phone_number,
  otp,
  otp_expiration_time
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users
WHERE phone_number = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CheckPhoneNumberUnique :one
SELECT COUNT(*) FROM users WHERE phone_number = $1;


-- name: UpdateUserByPhoneNumber :one
UPDATE users
set otp = $2,
otp_expiration_time = $3
WHERE phone_number = $1
RETURNING *;