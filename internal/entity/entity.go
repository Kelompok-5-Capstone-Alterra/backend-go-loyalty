package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Structs

type User struct {
	ID           uuid.UUID      `db:"id" gorm:"primaryKey;type:varchar(36)"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	Password     string         `db:"password"`
	MobileNumber string         `db:"mobile_number"`
	IsActive     bool           `db:"is_active"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	DeletedAt    gorm.DeletedAt `db:"deleted_at" gorm:"index"`
	RoleID       int            `db:"role_id"`
	Role         Role           `db:"role"`
	UserCoinID   uint64         `db:"user_coin_id"`
	UserCoin     UserCoin       `db:"user_coin"`
	CreditID     uint64         `db:"credit_id"`
	Credit       Credit         `db:"credit"`
}

type Role struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `db:"name"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
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
	ID            uint64         `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name          string         `db:"name"`
	Description   string         `db:"description"`
	RequiredPoint int64          `db:"required_points"`
	ValidUntil    time.Time      `db:"valid_until"`
	CategoryID    uint64         `db:"category_id"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
	DeletedAt     gorm.DeletedAt `db:"deleted_at" gorm:"index"`
	Category      Category       `db:"category"`
}

type Redeem struct {
	ID         uint64         `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	RewardID   uint64         `db:"reward_id"`
	UserID     uuid.UUID      `db:"user_id"`
	PointSpent int64          `db:"point_spent"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
	DeletedAt  gorm.DeletedAt `db:"deleted_at" gorm:"index"`
	Reward     Reward         `db:"reward"`
	User       User           `db:"user"`
}

type Redeems []Redeem

type Rewards []Reward

type Credit struct {
	ID     uint64 `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Amount int64  `db:"amount"`
}

type Credits []Credit

type FAQ struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement"`
	Question  string         `db:"question"`
	Answer    string         `db:"answer"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
}

type FAQs []FAQ

type Transaction struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID      `db:"user_id"`
	Status    string         `db:"status"`
	ProductID uint64         `db:"product_id"`
	Amount    int64          `db:"amount"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
	Product   Product        `db:"product"`
	User      User           `db:"user"`
}

type Transactions []Transaction

type Product struct {
	ID                 uint64         `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name               string         `db:"name"`
	Description        string         `db:"description"`
	Provider           string         `db:"provider"`
	ActivePeriod       int64          `db:"active_period"`
	Price              int64          `db:"price"`
	CategoryID         uint64         `db:"category_id"`
	MinimumTransaction uint32         `db:"minimum_transaction"`
	Coins              int            `db:"coins"`
	CreatedAt          time.Time      `db:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at"`
	DeletedAt          gorm.DeletedAt `db:"deleted_at" gorm:"index"`
	Category           Category       `db:"category"`
}

type ForgotPassword struct {
	ID        uint64    `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Email     string    `db:"email"`
	Token     string    `db:"token"`
	ExpiredAt time.Time `db:"expired_at"`
}

type Products []Product

type Category struct {
	ID        uint64         `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name      string         `db:"name"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"index"`
}

type Categories []Category

type PaymentInvoice struct {
	ID            uint64         `db:"id" gorm:"primaryKey;autoIncrement;column:id"`
	TransactionID uint64         `db:"transaction_id"`
	URL           string         `db:"url"`
	Amount        float64        `db:"amount"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
	DeletedAt     gorm.DeletedAt `db:"deleted_at" gorm:"index"`
	Transaction   Transaction    `gorm:"foreignKey:TransactionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type PaymentInvoices []PaymentInvoice

type UserCoin struct {
	ID uint64 `db:"id" gorm:"primaryKey;autoIncrement"`
	// UserID uuid.UUID `db:"user_id"`
	Amount int64 `db:"amount"`
	// User   User      `db:"user"`
}

type UserCoins []UserCoin
