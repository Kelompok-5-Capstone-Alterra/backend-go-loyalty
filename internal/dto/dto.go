package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Name         string `json:"name" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	MobileNumber string `json:"mobile_number" validate:"required"`
}
type UserUpdate struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
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
	ID           uuid.UUID    `json:"id"`
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
	Email string `json:"email" validate:"required,email"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type RewardRequest struct {
	Name          string `json:"name" validate:"required"`
	Description   string `json:"description" validate:"required"`
	RequiredPoint uint64 `json:"required_point" validate:"required"`
	ValidUntil    string `json:"valid_until" validate:"required"`
	CategoryID    uint64 `json:"category_id" validate:"required"`
}

type RewardUpdateRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	RequiredPoint uint64 `json:"required_point"`
	ValidUntil    string `json:"valid_until"`
	CategoryID    uint64 `json:"category_id"`
}

type RewardResponse struct {
	ID            uint64           `json:"id"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	RequiredPoint uint64           `json:"required_points"`
	ValidUntil    time.Time        `json:"valid_until"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     gorm.DeletedAt   `json:"deleted_at"`
	Category      CategoryResponse `json:"category"`
}

type RewardsResponse []RewardResponse
type ProductRequest struct {
	Name               string `json:"name" validate:"required"`
	CategoryID         uint64 `json:"category_id" validate:"required"`
	MinimumTransaction uint32 `json:"minimum_transaction" validate:"required"`
	Points             int    `json:"points" validate:"required"`
}
type ProductUpdateRequest struct {
	Name               string `json:"name"`
	CategoryID         uint64 `json:"category_id"`
	MinimumTransaction uint32 `json:"minimum_transaction"`
	Points             int    `json:"points"`
}

type CategoryResponse struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CategoriesResponse []CategoryResponse

type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type ForgotPasswordTokenRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type NewPassword struct {
	Password string `json:"password" validate:"required"`
}
type ProductResponse struct {
	ID                 uint64           `json:"id"`
	Name               string           `json:"name"`
	CategoryID         uint64           `json:"category_id"`
	MinimumTransaction uint32           `json:"minimum_transaction"`
	Points             int              `json:"points"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	DeletedAt          gorm.DeletedAt   `json:"deleted_at"`
	Category           CategoryResponse `json:"category"`
}

type ProductsResponse []ProductResponse

type RedeemRequest struct {
	RewardID uint64 `json:"reward_id" validate:"required"`
}
type RedeemUpdateRequest struct {
	RewardID   uint64 `json:"reward_id"`
	PointSpent uint64 `json:"point_spent"`
}

type RedeemResponse struct {
	ID         uint64         `json:"id"`
	PointSpent uint64         `json:"point_spent"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Reward     RewardResponse `json:"reward"`
}

type RedeemResponses []RedeemResponse

type UserCoinResponse struct {
	ID     uint64 `json:"id"`
	Amount uint64 `json:"amount"`
	// User   UserResponse `json:"user"`
}

type UserCoinResponses []UserCoinResponse

type CreditResponse struct {
	ID     uint64 `json:"id"`
	Amount uint64 `json:"amount"`
}

type CreditResponses []CreditResponse
