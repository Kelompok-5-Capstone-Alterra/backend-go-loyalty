package dto

import (
	"time"
)

type SignUpRequest struct {
	Name         string `json:"name" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" validate:"required"`
	MobileNumber string `json:"mobile_number" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserResponse struct {
	ID           uint64       `json:"id"`
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	MobileNumber string       `json:"mobile_number"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Role         RoleResponse `json:"role"`
}

type RoleResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponses []UserResponse
type RoleResponses []RoleResponse

type JWTData struct {
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	MobileNumber string       `json:"mobile_number"`
	Role         RoleResponse `json:"role"`
}

type ValidateOTP struct {
	OTP   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type RequestNewOTP struct {
	Email string `json:"email" validate:"required"`
}
