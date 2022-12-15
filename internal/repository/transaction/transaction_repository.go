package transactionRepository

import (
	"backend-go-loyalty/internal/entity"
	"context"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetAllTransaction(ctx context.Context) (entity.Transactions, error)
	GetTransactionByID(ctx context.Context, id uint64) (entity.Transaction, error)
	InsertTransaction(ctx context.Context, req entity.Transaction) error
	UpdateTransaction(ctx context.Context, req entity.Transaction) error
	DeleteTransaction(ctx context.Context, id uint64) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionRepository {
	return transactionRepository{
		db: db,
	}
}

func (tr transactionRepository) GetAllTransaction(ctx context.Context) (entity.Transactions, error)
func (tr transactionRepository) GetTransactionByID(ctx context.Context, id uint64) (entity.Transaction, error)
func (tr transactionRepository) InsertTransaction(ctx context.Context, req entity.Transaction) error
func (tr transactionRepository) UpdateTransaction(ctx context.Context, req entity.Transaction) error
func (tr transactionRepository) DeleteTransaction(ctx context.Context, id uint64) error
