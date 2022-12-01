package entity

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
	ID            uint64       `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name          string       `db:"name"`
	Description   string       `db:"description"`
	RequiredPoint uint64       `db:"required_points"`
	ValidUntil    time.Time    `db:"valid_until"`
	CategoryID    uint64       `db:"category_id"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
	Category      Category     `db:"category"`
}

type Redeem struct {
	ID         uint64       `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	RewardID   uint64       `db:"reward_id"`
	UserID     uint64       `db:"user_id"`
	PointSpent uint64       `db:"point_spent"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
}

type Rewards []Reward

type Credit struct {
	ID     uint64 `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserID uint64 `db:"user_id"`
	Amount uint64 `db:"amount"`
}

type Point struct {
	ID     uint64 `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserID uint64 `db:"user_id"`
	Amount uint64 `db:"amount"`
}

type Credits []Credit
type Points []Point

type FAQ struct {
	ID        uint64       `db:"id" gorm:"primaryKey;autoIncrement"`
	Question  string       `db:"question"`
	Answer    string       `db:"answer"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type FAQs []FAQ

type Transaction struct {
	ID        uint64       `db:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint64       `db:"user_id"`
	Status    string       `db:"status"`
	Amount    uint64       `db:"amount"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Transactions []Transaction

type TransactionDetail struct {
	ID            uint64       `db:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID uint64       `db:"transaction_id"`
	ProductID     uint64       `db:"product_id"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}

type TransactionDetails []TransactionDetail

type Product struct {
	ID                 uint64       `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name               string       `db:"name"`
	CategoryID         uint64       `db:"category_id"`
	MinimumTransaction uint32       `db:"minimum_transaction"`
	Points             int          `db:"points"`
	CreatedAt          time.Time    `db:"created_at"`
	UpdatedAt          time.Time    `db:"updated_at"`
	DeletedAt          sql.NullTime `db:"deleted_at"`
	Category           Category     `db:"category"`
}

type Products []Product

type Category struct {
	ID        uint64       `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Categories []Category

type PaymentInvoice struct {
	ID            uint64       `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	TransactionID uint64       `db:"transaction_id"`
	URL           string       `db:"url"`
	Amount        uint64       `db:"amount"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}

type PaymentInvoices []PaymentInvoice
