package model

import (
	"database/sql"
	"time"
)

// Structs

type User struct {
	ID           uint64       `db:"id" gorm:"primaryKey;autoIncrement"`
	Name         string       `db:"name"`
	Email        string       `db:"email"`
	Password     string       `db:"password"`
	MobileNumber string       `db:"mobile_number"`
	IsActive     bool         `db:"is_active"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at" gorm:"index"`
	RoleID       int          `db:"role_id"`
	Role         Role         `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type Product struct {
	ProductID   uint64    `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name        string    `db:"name"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
	Point       int       `db:"points"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
