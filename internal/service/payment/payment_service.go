package paymentService

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/internal/entity"
	paymentRepository "backend-go-loyalty/internal/repository/payment"
	transactionRepository "backend-go-loyalty/internal/repository/transaction"
	userRepository "backend-go-loyalty/internal/repository/user"
	"backend-go-loyalty/pkg/utils"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

type IPaymentService interface {
	PayWithCredit(ctx context.Context, userID uuid.UUID, transactionID uint64) error
	PayWithOVO(ctx context.Context, req dto.PayWithOVO, userID uuid.UUID) (*xendit.EWalletCharge, error)
}

type paymentService struct {
	tr transactionRepository.ITransactionRepository
	pr paymentRepository.IPaymentRepository
	ur userRepository.UserRepositoryInterface
}

func NewPaymentService(tr transactionRepository.ITransactionRepository, pr paymentRepository.IPaymentRepository, ur userRepository.UserRepositoryInterface) paymentService {
	return paymentService{
		tr: tr,
		pr: pr,
		ur: ur,
	}
}

func (ps paymentService) PayWithOVO(ctx context.Context, req dto.PayWithOVO, userID uuid.UUID) (*xendit.EWalletCharge, error) {
	// Get Transaction Data
	transaction, err := ps.tr.GetTransactionByIDByUserID(ctx, req.TransactionID, userID)
	if err != nil {
		return nil, err
	}
	if transaction.Status == "SUCCEEDED" || transaction.Status == "REFUNDED" || transaction.Status == "FAILED" {
		return nil, errors.New("cannot charge")
	}

	params := ewallet.CreateEWalletChargeParams{
		ReferenceID:    utils.CreateExternalID(req.TransactionID),
		Currency:       "IDR",
		Amount:         transaction.Amount,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_OVO",
		ChannelProperties: map[string]string{
			"mobile_number": req.MobileNumber,
		},
		Metadata: map[string]interface{}{
			"product_name": transaction.Product.Name,
			"ordered_at":   transaction.CreatedAt,
		},
	}
	res, err := ps.pr.CreateEWalletCharge(ctx, &params)
	return res, err
}

func (ps paymentService) PayWithCredit(ctx context.Context, userID uuid.UUID, transactionID uint64) error {
	// Get user data
	user, err := ps.ur.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Get Transaction Data
	transaction, err := ps.tr.GetTransactionByIDByUserID(ctx, transactionID, userID)
	if err != nil {
		return err
	}

	sub := user.Credit.Amount - transaction.Amount
	if sub < 0 {
		return errors.New("not enough credit")
	}

	// Updates Transaction Status
	update := entity.Transaction{
		Status: "SUCCESS",
	}
	err = ps.tr.UpdateTransaction(ctx, update, transactionID)
	if err != nil {
		return err
	}

	// Updates Credit Value
	updateCredit := entity.Credit{
		ID:     user.CreditID,
		Amount: sub,
	}
	err = ps.pr.UpdateCredit(ctx, updateCredit)
	return err
}
