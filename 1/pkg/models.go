package pkg

import "time"

const OtpValidDuration = 1 * time.Minute

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=255"`
	PhoneNumber string `json:"phone_number" binding:"required,min=1,max=30"`
}

type GenerateOtpRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,min=1,max=30"`
}

type VerifyOtpRequest struct {
	OTP         string `json:"otp" binding:"required,min=4,max=4"`
	PhoneNumber string `json:"phone_number" binding:"required,min=1,max=30"`
}

type User struct {
	ID          int32  `json:"ID"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
