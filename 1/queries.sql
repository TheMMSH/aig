-- name: CreateUser :one
INSERT INTO users (name, phone_number)
VALUES ($1, $2)
RETURNING *;

-- name: GenerateOTP :one
UPDATE users
SET otp = $1, otp_expiration_time = $2
WHERE phone_number = $3
RETURNING *;

-- name: VerifyOTP :one
SELECT *
FROM users
WHERE phone_number = $1;
