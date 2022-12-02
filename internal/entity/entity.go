package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Structs

type User struct {
	ID           uuid.UUID    `db:"id" gorm:"primaryKey;type:varchar(36)"`
	Name         string       `db:"name"`
	Email        string       `db:"email"`
	Password     string       `db:"password"`
	MobileNumber string       `db:"mobile_number"`
	IsActive     bool         `db:"is_active"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at" gorm:"index"`
	RoleID       int          `db:"role_id"`
	Role         Role         `db:"role"`
}

type Role struct {
	ID        uint64       `db:"id" gorm:"primaryKey;autoIncrement"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" gorm:"index"`
}

type OTP struct {
	ID        uint64    `db:"id" gorm:"primaryKey;autoIncrement"`
	OTPCode   string    `db:"otp_code"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

// Array of structs
type Users []User
type Roles []Role

type Reward struct {
	RewardID      uint64 `db:"id" gorm:"column:id"`
	Name          string `db:"name"`
	Description   string `db:"description"`
	RequiredPoint uint64 `db:"required_points"`
}

type Rewards []Reward
type Product struct {
	ProductID   uint64 `db:"id" gorm:"column:id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	Description string `db:"description"`
	Point       int    `db:"points"`
}

type Products []Product
