package transactionService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	paymentRepository "backend-go-loyalty/internal/repository/payment"
	productRepository "backend-go-loyalty/internal/repository/product"
	transactionRepository "backend-go-loyalty/internal/repository/transaction"
	userRepository "backend-go-loyalty/internal/repository/user"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ITransactionService interface {
	GetTransactionByUserID(ctx context.Context, id uuid.UUID) (dto.TransactionResponses, error)
	GetTransactionByIDByUserID(ctx context.Context, user_id uuid.UUID, id uint64) (dto.TransactionResponse, error)
	GetAllTransaction(ctx context.Context) (dto.TransactionResponses, error)
	GetTransactionByID(ctx context.Context, id uint64) (dto.TransactionResponse, error)
	CreateTransaction(ctx context.Context, req dto.TransactionRequest, id uuid.UUID) error
	UpdateStatus(ctx context.Context, status string, id uint64) error
	DeleteTransaction(ctx context.Context, id uint64) error
}

type transactionService struct {
	tr  transactionRepository.ITransactionRepository
	ur  userRepository.UserRepositoryInterface
	pr  productRepository.IProductRepository
	pyr paymentRepository.IPaymentRepository
}

func NewTransactionService(tr transactionRepository.ITransactionRepository, ur userRepository.UserRepositoryInterface, pr productRepository.IProductRepository, pyr paymentRepository.IPaymentRepository) transactionService {
	return transactionService{
		tr:  tr,
		ur:  ur,
		pr:  pr,
		pyr: pyr,
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
			User: dto.UserResponse{
				ID:           val.User.ID,
				Name:         val.User.Name,
				Email:        val.User.Email,
				MobileNumber: val.User.MobileNumber,
				CreatedAt:    val.User.CreatedAt,
				UpdatedAt:    val.User.UpdatedAt,
				Role: dto.RoleResponse{
					ID:        val.User.Role.ID,
					Name:      val.User.Role.Name,
					CreatedAt: val.User.Role.CreatedAt,
					UpdatedAt: val.User.Role.UpdatedAt,
				},
				UserCoin: dto.UserCoinResponse{
					ID:     val.User.UserCoin.ID,
					Amount: val.User.UserCoin.Amount,
				},
				Credit: dto.CreditResponse{
					ID:     val.User.Credit.ID,
					Amount: val.User.Credit.Amount,
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
		User: dto.UserResponse{
			ID:           data.User.ID,
			Name:         data.User.Name,
			Email:        data.User.Email,
			MobileNumber: data.User.MobileNumber,
			CreatedAt:    data.User.CreatedAt,
			UpdatedAt:    data.User.UpdatedAt,
			Role: dto.RoleResponse{
				ID:        data.User.Role.ID,
				Name:      data.User.Role.Name,
				CreatedAt: data.User.Role.CreatedAt,
				UpdatedAt: data.User.Role.UpdatedAt,
			},
			UserCoin: dto.UserCoinResponse{
				ID:     data.User.UserCoin.ID,
				Amount: data.User.UserCoin.Amount,
			},
			Credit: dto.CreditResponse{
				ID:     data.User.Credit.ID,
				Amount: data.User.Credit.Amount,
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
			User: dto.UserResponse{
				ID:           val.User.ID,
				Name:         val.User.Name,
				Email:        val.User.Email,
				MobileNumber: val.User.MobileNumber,
				CreatedAt:    val.User.CreatedAt,
				UpdatedAt:    val.User.UpdatedAt,
				Role: dto.RoleResponse{
					ID:        val.User.Role.ID,
					Name:      val.User.Role.Name,
					CreatedAt: val.User.Role.CreatedAt,
					UpdatedAt: val.User.Role.UpdatedAt,
				},
				UserCoin: dto.UserCoinResponse{
					ID:     val.User.UserCoin.ID,
					Amount: val.User.UserCoin.Amount,
				},
				Credit: dto.CreditResponse{
					ID:     val.User.Credit.ID,
					Amount: val.User.Credit.Amount,
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
		User: dto.UserResponse{
			ID:           data.User.ID,
			Name:         data.User.Name,
			Email:        data.User.Email,
			MobileNumber: data.User.MobileNumber,
			CreatedAt:    data.User.CreatedAt,
			UpdatedAt:    data.User.UpdatedAt,
			Role: dto.RoleResponse{
				ID:        data.User.Role.ID,
				Name:      data.User.Role.Name,
				CreatedAt: data.User.Role.CreatedAt,
				UpdatedAt: data.User.Role.UpdatedAt,
			},
			UserCoin: dto.UserCoinResponse{
				ID:     data.User.UserCoin.ID,
				Amount: data.User.UserCoin.Amount,
			},
			Credit: dto.CreditResponse{
				ID:     data.User.Credit.ID,
				Amount: data.User.Credit.Amount,
			},
		},
	}
	return transaction, err
}
func (ts transactionService) CreateTransaction(ctx context.Context, req dto.TransactionRequest, id uuid.UUID) error {
	user, err := ts.ur.GetUserByID(ctx, id)
	if err != nil {
		errorMsg := "[user] " + err.Error()
		return errors.New(errorMsg)
	}
	product, err := ts.pr.GetProductByID(ctx, req.ProductID)
	if err != nil {
		errorMsg := "[product] " + err.Error()
		return errors.New(errorMsg)
	}
	sub := user.Credit.Amount - product.Price
	if sub < 0 {
		return errors.New("not enough credit")
	}
	transactionRequest := entity.Transaction{
		UserID:    id,
		Status:    "PENDING",
		ProductID: req.ProductID,
		Amount:    product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	transaction, err := ts.tr.InsertTransaction(ctx, transactionRequest)
	if err != nil {
		errorMsg := "[transaction] " + err.Error()
		return errors.New(errorMsg)
	}
	invo, err := ts.pyr.CreateInvoice(ctx, transaction, user)
	if err != nil {
		return err
	}
	err = ts.pyr.InsertInvoiceData(ctx, invo, transaction.ID)
	return err
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
