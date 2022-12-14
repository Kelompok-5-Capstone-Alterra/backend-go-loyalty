package transactionService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	paymentRepository "backend-go-loyalty/internal/repository/payment"
	pointRepository "backend-go-loyalty/internal/repository/point"
	productRepository "backend-go-loyalty/internal/repository/product"
	transactionRepository "backend-go-loyalty/internal/repository/transaction"
	userRepository "backend-go-loyalty/internal/repository/user"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ITransactionService interface {
	GetTransactionByUserID(ctx context.Context, id uuid.UUID) (dto.TransactionResponses, error)
	GetTransactionByIDByUserID(ctx context.Context, user_id uuid.UUID, id uint64) (dto.TransactionResponse, error)
	GetAllTransaction(ctx context.Context) (dto.TransactionResponses, error)
	GetTransactionByID(ctx context.Context, id uint64) (dto.TransactionResponse, error)
	CreateTransaction(ctx context.Context, req dto.TransactionRequest, id uuid.UUID) (dto.TransactionResponse,error)
	UpdateStatus(ctx context.Context, status string, id uint64) error
	DeleteTransaction(ctx context.Context, id uint64) error
	CheckCoinEligibility(ctx context.Context, userID uuid.UUID, transactionID uint64) error
	DeleteInvoice(ctx context.Context, transactionID uint64) error
}

type transactionService struct {
	tr  transactionRepository.ITransactionRepository
	ur  userRepository.UserRepositoryInterface
	pr  productRepository.IProductRepository
	pyr paymentRepository.IPaymentRepository
	cr  pointRepository.IPointRepository
}

func NewTransactionService(tr transactionRepository.ITransactionRepository, ur userRepository.UserRepositoryInterface, pr productRepository.IProductRepository, pyr paymentRepository.IPaymentRepository, cr pointRepository.IPointRepository) transactionService {
	return transactionService{
		tr:  tr,
		ur:  ur,
		pr:  pr,
		pyr: pyr,
		cr:  cr,
	}
}

func (ts transactionService) GetTransactionByUserID(ctx context.Context, id uuid.UUID) (dto.TransactionResponses, error) {
	data, err := ts.tr.GetTransactionByUserID(ctx, id)
	if err != nil {
		return dto.TransactionResponses{}, err
	}
	var transactions dto.TransactionResponses
	for _, val := range data {
		transaction := dto.TransactionResponse{
			ID:        val.ID,
			UserID:    val.UserID,
			Status:    val.Status,
			Amount:    val.Amount,
			ProductID: val.ProductID,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			DeletedAt: val.DeletedAt,
			Product: dto.ProductResponse{
				ID:                 val.Product.ID,
				Name:               val.Product.Name,
				Description:        val.Product.Description,
				Provider:           val.Product.Provider,
				ActivePeriod:       val.Product.ActivePeriod,
				Price:              val.Product.Price,
				MinimumTransaction: val.Product.MinimumTransaction,
				Coins:              val.Product.Coins,
				CreatedAt:          val.Product.CreatedAt,
				UpdatedAt:          val.Product.UpdatedAt,
				DeletedAt:          val.Product.DeletedAt,
				Category: dto.CategoryResponse{
					ID:        val.Product.Category.ID,
					Name:      val.Product.Category.Name,
					CreatedAt: val.Product.Category.CreatedAt,
					UpdatedAt: val.Product.Category.UpdatedAt,
					DeletedAt: val.Product.Category.DeletedAt,
				},
			},
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err
}
func (ts transactionService) GetTransactionByIDByUserID(ctx context.Context, user_id uuid.UUID, id uint64) (dto.TransactionResponse, error) {
	data, err := ts.tr.GetTransactionByIDByUserID(ctx, id, user_id)
	transaction := dto.TransactionResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		Status:    data.Status,
		Amount:    data.Amount,
		ProductID: data.ProductID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Product: dto.ProductResponse{
			ID:                 data.Product.ID,
			Name:               data.Product.Name,
			Description:        data.Product.Description,
			Provider:           data.Product.Provider,
			ActivePeriod:       data.Product.ActivePeriod,
			Price:              data.Product.Price,
			MinimumTransaction: data.Product.MinimumTransaction,
			Coins:              data.Product.Coins,
			CreatedAt:          data.Product.CreatedAt,
			UpdatedAt:          data.Product.UpdatedAt,
			DeletedAt:          data.Product.DeletedAt,
			Category: dto.CategoryResponse{
				ID:        data.Product.Category.ID,
				Name:      data.Product.Category.Name,
				CreatedAt: data.Product.Category.CreatedAt,
				UpdatedAt: data.Product.Category.UpdatedAt,
				DeletedAt: data.Product.Category.DeletedAt,
			},
		},
	}
	return transaction, err
}

func (ts transactionService) GetAllTransaction(ctx context.Context) (dto.TransactionResponses, error) {
	data, err := ts.tr.GetAllTransaction(ctx)
	if err != nil {
		return dto.TransactionResponses{}, err
	}
	var transactions dto.TransactionResponses
	for _, val := range data {
		transaction := dto.TransactionResponse{
			ID:        val.ID,
			UserID:    val.UserID,
			Status:    val.Status,
			Amount:    val.Amount,
			ProductID: val.ProductID,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			DeletedAt: val.DeletedAt,
			Product: dto.ProductResponse{
				ID:                 val.Product.ID,
				Name:               val.Product.Name,
				Description:        val.Product.Description,
				Provider:           val.Product.Provider,
				ActivePeriod:       val.Product.ActivePeriod,
				Price:              val.Product.Price,
				MinimumTransaction: val.Product.MinimumTransaction,
				Coins:              val.Product.Coins,
				CreatedAt:          val.Product.CreatedAt,
				UpdatedAt:          val.Product.UpdatedAt,
				DeletedAt:          val.Product.DeletedAt,
				Category: dto.CategoryResponse{
					ID:        val.Product.Category.ID,
					Name:      val.Product.Category.Name,
					CreatedAt: val.Product.Category.CreatedAt,
					UpdatedAt: val.Product.Category.UpdatedAt,
					DeletedAt: val.Product.Category.DeletedAt,
				},
			},
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err
}
func (ts transactionService) GetTransactionByID(ctx context.Context, id uint64) (dto.TransactionResponse, error) {
	data, err := ts.tr.GetTransactionByID(ctx, id)
	transaction := dto.TransactionResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		Status:    data.Status,
		Amount:    data.Amount,
		ProductID: data.ProductID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Product: dto.ProductResponse{
			ID:                 data.Product.ID,
			Name:               data.Product.Name,
			Description:        data.Product.Description,
			Provider:           data.Product.Provider,
			ActivePeriod:       data.Product.ActivePeriod,
			Price:              data.Product.Price,
			MinimumTransaction: data.Product.MinimumTransaction,
			Coins:              data.Product.Coins,
			CreatedAt:          data.Product.CreatedAt,
			UpdatedAt:          data.Product.UpdatedAt,
			DeletedAt:          data.Product.DeletedAt,
			Category: dto.CategoryResponse{
				ID:        data.Product.Category.ID,
				Name:      data.Product.Category.Name,
				CreatedAt: data.Product.Category.CreatedAt,
				UpdatedAt: data.Product.Category.UpdatedAt,
				DeletedAt: data.Product.Category.DeletedAt,
			},
		},
	}
	return transaction, err
}
func (ts transactionService) CreateTransaction(ctx context.Context, req dto.TransactionRequest, id uuid.UUID) (dto.TransactionResponse,error) {
	// Get user data
	// user, err := ts.ur.GetUserByID(ctx, id)
	// if err != nil {
	// 	errorMsg := "[user] " + err.Error()
	// 	return nil, errors.New(errorMsg)
	// }
	// Get product data
	product, err := ts.pr.GetProductByID(ctx, req.ProductID)
	if err != nil {
		errorMsg := "[product] " + err.Error()
		return dto.TransactionResponse{},errors.New(errorMsg)
	}
	// // Check if credit is enough
	// sub := user.Credit.Amount - product.Price
	// if sub < 0 {
	// 	return errors.New("not enough credit")
	// }

	// Create transaction request
	transactionRequest := entity.Transaction{
		UserID:    id,
		Status:    "PENDING",
		ProductID: req.ProductID,
		Amount:    product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert transaction to database
	data, err := ts.tr.InsertTransaction(ctx, transactionRequest)
	transaction := dto.TransactionResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		Status:    data.Status,
		Amount:    data.Amount,
		ProductID: data.ProductID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		Product: dto.ProductResponse{
			ID:                 data.Product.ID,
			Name:               data.Product.Name,
			Description:        data.Product.Description,
			Provider:           data.Product.Provider,
			ActivePeriod:       data.Product.ActivePeriod,
			Price:              data.Product.Price,
			MinimumTransaction: data.Product.MinimumTransaction,
			Coins:              data.Product.Coins,
			CreatedAt:          data.Product.CreatedAt,
			UpdatedAt:          data.Product.UpdatedAt,
			DeletedAt:          data.Product.DeletedAt,
			Category: dto.CategoryResponse{
				ID:        data.Product.Category.ID,
				Name:      data.Product.Category.Name,
				CreatedAt: data.Product.Category.CreatedAt,
				UpdatedAt: data.Product.Category.UpdatedAt,
				DeletedAt: data.Product.Category.DeletedAt,
			},
		},
	}
	return transaction, err

	// Create xendit invoice
	// invo, err := ts.pyr.CreateInvoice(ctx, transaction, user)
	// if err != nil {
	// 	return nil, err
	// }

	// // Insert invoice data to database
	// err = ts.pyr.InsertInvoiceData(ctx, invo, transaction.ID)

	// // Return invoice data and errors
}

func (ts transactionService) UpdateStatus(ctx context.Context, status string, id uint64) error {
	req := entity.Transaction{
		Status: status,
	}
	err := ts.tr.UpdateTransaction(ctx, req, id)
	return err
}

func (ts transactionService) DeleteTransaction(ctx context.Context, id uint64) error {
	err := ts.tr.DeleteTransaction(ctx, id)
	return err
}
func (ts transactionService) CheckCoinEligibility(ctx context.Context, userID uuid.UUID, transactionID uint64) error {
	transaction, err := ts.tr.GetTransactionByID(ctx, transactionID)
	count, err := ts.tr.CountSuccessTransactionByUserID(ctx, userID)
	fmt.Println(count, transaction.Product.MinimumTransaction)
	if err != nil {
		return err
	}
	if count%int64(transaction.Product.MinimumTransaction) == 0 {
		req := entity.Transaction{
			CoinsEarned: int64(transaction.Product.Coins),
		}
		err = ts.tr.UpdateTransaction(ctx, req, transaction.ID)
		if err != nil {
			return err
		}
		user, err := ts.ur.GetUserByID(ctx, userID)
		if err != nil {
			return err
		}
		userCoin := entity.UserCoin{
			Amount: user.UserCoin.Amount + int64(transaction.Product.Coins),
		}
		err = ts.cr.UpdatePoint(ctx, user.UserCoinID, userCoin)
		return err
	}
	return nil
}

func (ts transactionService) DeleteInvoice(ctx context.Context, transactionID uint64) error {
	err := ts.tr.DeleteInvoice(ctx, transactionID)
	return err
}
