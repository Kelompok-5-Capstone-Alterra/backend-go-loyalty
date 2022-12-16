package transactionRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetAllTransaction(ctx context.Context) (entity.Transactions, error)
	GetTransactionByIDByUserID(ctx context.Context, id uint64, user_id uuid.UUID) (entity.Transaction, error)
	GetTransactionByID(ctx context.Context, id uint64) (entity.Transaction, error)
	GetTransactionByUserID(ctx context.Context, id uuid.UUID) (entity.Transactions, error)
	InsertTransaction(ctx context.Context, req entity.Transaction) (entity.Transaction, error)
	UpdateTransaction(ctx context.Context, req entity.Transaction, id uint64) error
	DeleteTransaction(ctx context.Context, id uint64) error
	CountSuccessTransactionByUserID(ctx context.Context, userID uuid.UUID)(int64, error)
	DeleteInvoice(ctx context.Context, transactionID uint64) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionRepository {
	return transactionRepository{
		db: db,
	}
}

func (tr transactionRepository) GetTransactionByIDByUserID(ctx context.Context, id uint64, user_id uuid.UUID) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := tr.db.Model(&model.Transaction{}).Preload("Product").Preload("User").Where("user_id = ?", user_id).First(&transaction, id).Error
	return transaction, err
}

func (tr transactionRepository) GetTransactionByUserID(ctx context.Context, id uuid.UUID) (entity.Transactions, error) {
	var transactions entity.Transactions
	err := tr.db.Model(&model.Transaction{}).Preload("Product").Preload("User").Where("user_id = ?", id).Find(&transactions).Error
	return transactions, err
}

func (tr transactionRepository) GetAllTransaction(ctx context.Context) (entity.Transactions, error) {
	var transactions entity.Transactions
	err := tr.db.Model(&model.Transaction{}).Preload("Product").Preload("User").Find(&transactions).Error
	return transactions, err
}
func (tr transactionRepository) GetTransactionByID(ctx context.Context, id uint64) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := tr.db.Model(&model.Transaction{}).Preload("Product").Preload("User").First(&transaction, id).Error
	return transaction, err
}
func (tr transactionRepository) InsertTransaction(ctx context.Context, req entity.Transaction) (entity.Transaction, error) {
	err := tr.db.Create(&req).Error
	return req, err
}

func (tr transactionRepository) UpdateTransaction(ctx context.Context, req entity.Transaction, id uint64) error {
	err := tr.db.Model(&model.Transaction{}).Where("id = ?", id).Updates(req).Error
	return err
}

func (tr transactionRepository) DeleteTransaction(ctx context.Context, id uint64) error {
	err := tr.db.Delete(&model.Transaction{}, id).Error
	return err
}

func (tr transactionRepository) CountSuccessTransactionByUserID(ctx context.Context, userID uuid.UUID)(int64, error){
	var count int64
	err := tr.db.Model(&model.Transaction{}).Where("user_id = ? AND status = ", userID, "SUCCESS").Count(&count).Error
	return count, err
}

func (tr transactionRepository) DeleteInvoice(ctx context.Context, transactionID uint64) error{
	err := tr.db.Where("transaction_id = ?",transactionID).Delete(&model.PaymentInvoice{}).Error
	return err
}