// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, phone_number)
VALUES ($1, $2)
RETURNING id, name, phone_number, otp, otp_expiration_time
`

type CreateUserParams struct {
	Name        pgtype.Text
	PhoneNumber string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, name, phone_number, otp, otp_expiration_time
FROM users
WHERE phone_number = $1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phoneNumber string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhone, phoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}

const updateUserOTP = `-- name: UpdateUserOTP :one
UPDATE users
SET otp = $1, otp_expiration_time = $2
WHERE id = $3
RETURNING id, name, phone_number, otp, otp_expiration_time
`

type UpdateUserOTPParams struct {
	Otp               pgtype.Text
	OtpExpirationTime pgtype.Timestamp
	ID                int32
}

func (q *Queries) UpdateUserOTP(ctx context.Context, arg UpdateUserOTPParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserOTP, arg.Otp, arg.OtpExpirationTime, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
	)
	return i, err
}
