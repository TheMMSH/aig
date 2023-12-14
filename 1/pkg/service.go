package pkg

import (
	"aig/1/repository"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Service struct {
	repo *repository.Queries
	DB   *pgx.Conn
}

func NewService(db *pgx.Conn) Service {
	return Service{
		repo: repository.New(db),
		DB:   db,
	}
}

func (s Service) CreateUser(ctx context.Context, user CreateUserRequest) (*User, error) {
	u, err := s.repo.CreateUser(ctx, toCreateUserParam(user))
	if err != nil {
		return nil, err
	}

	return toUser(u), err
}

func (s Service) GenerateOtp(ctx context.Context, phoneNumber string) (time.Time, error) {
	expTime := time.Now().Add(OtpValidDuration)
	randToken, err := generateRandomSequence(4)
	if err != nil {
		return time.Now(), errors.New("cannot generate random token, try again later")
	}

	_, err = s.repo.GenerateOTP(ctx, repository.GenerateOTPParams{
		Otp: pgtype.Text{
			String: string(randToken),
			Valid:  true,
		},
		OtpExpirationTime: pgtype.Timestamp{
			Time:  expTime,
			Valid: true,
		},
		PhoneNumber: phoneNumber,
	})

	return expTime, err
}

func (s Service) VerifyOTP(ctx context.Context, req VerifyOtpRequest) error {
	user, err := s.repo.VerifyOTP(ctx, req.PhoneNumber)

	if err != nil {
		return err
	}

	if user.Otp.String != req.OTP {
		return errors.New("error invalid otp")
	}

	if user.OtpExpirationTime.Time.Before(time.Now()) {
		return errors.New("otp expired")
	}

	return nil
}

func toUser(u repository.User) *User {
	return &User{
		ID:          u.ID,
		Name:        u.Name.String,
		PhoneNumber: u.PhoneNumber,
	}
}

func toCreateUserParam(user CreateUserRequest) repository.CreateUserParams {
	return repository.CreateUserParams{
		Name: pgtype.Text{
			String: user.Name,
			Valid:  true,
		},
		PhoneNumber: user.PhoneNumber,
	}
}
