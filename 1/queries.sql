-- name: CreateUser :one
INSERT INTO users (name, phone_number)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUserOTP :one
UPDATE users
SET otp = $1, otp_expiration_time = $2
WHERE id = $3
RETURNING *;

-- name: GetUserByPhone :one
SELECT *
FROM users
WHERE phone_number = $1;
