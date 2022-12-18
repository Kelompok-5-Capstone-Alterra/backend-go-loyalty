package paymentRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/interfaces"
	"backend-go-loyalty/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	// CreateInvoice(ctx context.Context, req entity.Transaction, user entity.User) (*xendit.Invoice, error)
	InsertInvoiceData(ctx context.Context, req *xendit.Invoice, transactionID uint64) error
	// DeleteInvoiceByTransactionID(ctx context.Context, transactionID uint64) error
	UpdateCredit(ctx context.Context, req entity.Credit) error
	CreateEWalletCharge(ctx context.Context, data *ewallet.CreateEWalletChargeParams) (*xendit.EWalletCharge, error)
}

type paymentRepository struct {
	EWallet interfaces.XenditEWallet
	db      *gorm.DB
}

func NewPaymentRepository(db *gorm.DB, xenditEwallet interfaces.XenditEWallet) paymentRepository {
	return paymentRepository{
		EWallet: xenditEwallet,
		db:      db,
	}
}

func (pr paymentRepository) UpdateCredit(ctx context.Context, req entity.Credit) error {
	err := pr.db.Model(&model.Credit{}).Where("id = ?", req.ID).Updates(req).Error
	return err
}

func (pr paymentRepository) CreateEWalletCharge(ctx context.Context, data *ewallet.CreateEWalletChargeParams) (*xendit.EWalletCharge, error) {
	charge, err := pr.EWallet.CreateEWalletCharge(data)
	if err != nil{
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}
	return charge, nil
}

// func (pr paymentRepository) CreateInvoice(ctx context.Context, req entity.Transaction, user entity.User) (*xendit.Invoice, error) {
// 	customer := xendit.InvoiceCustomer{
// 		GivenNames:   user.Name,
// 		Email:        user.Email,
// 		MobileNumber: user.MobileNumber,
// 		Address:      nil,
// 	}
// 	items := []xendit.InvoiceItem{
// 		{
// 			Name:     req.Product.Name,
// 			Price:    float64(req.Product.Price),
// 			Quantity: 1,
// 		},
// 	}
// 	fee := xendit.InvoiceFee{
// 		Type:  "ADMIN",
// 		Value: 5000,
// 	}
// 	NotificationType := []string{"whatsapp", "email"}

// 	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
// 		InvoiceCreated:  NotificationType,
// 		InvoiceReminder: NotificationType,
// 		InvoicePaid:     NotificationType,
// 		InvoiceExpired:  NotificationType,
// 	}

// 	invoiceParams := invoice.CreateParams{
// 		ExternalID:                     utils.CreateExternalID(req.ID),
// 		Amount:                         float64(req.Amount + 5000),
// 		Description:                    "Invoice Pemesanan",
// 		InvoiceDuration:                3600,
// 		Customer:                       customer,
// 		CustomerNotificationPreference: customerNotificationPreference,
// 		SuccessRedirectURL:             "https://www.google.com",
// 		FailureRedirectURL:             "https://www.google.com",
// 		Currency:                       "IDR",
// 		Items:                          items,
// 		Fees:                           []xendit.InvoiceFee{fee},
// 	}
// invo, err := pr.Xendit.Create(&invoiceParams)
// 	return invo, err
// }
func (pr paymentRepository) InsertInvoiceData(ctx context.Context, req *xendit.Invoice, transactionID uint64) error {
	invoice := entity.PaymentInvoice{
		InvoiceID:     req.ID,
		TransactionID: transactionID,
		URL:           req.InvoiceURL,
		Amount:        req.Amount,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := pr.db.Create(&invoice).Error
	return err
}

// func (pr paymentRepository) DeleteInvoiceByTransactionID(ctx context.Context, transactionID uint64) error {
// 	payment := entity.PaymentInvoice{}
// 	err := pr.db.Model(&model.PaymentInvoice{}).Where("transaction_id = ?", transactionID).First(&payment).Error
// 	if err != nil {
// 		return err
// 	}
// 	err = pr.db.Where("transaction_id = ?", transactionID).Delete(&entity.PaymentInvoice{}).Error
// 	if err != nil {
// 		return err
// 	}
// 	cancelInvoice := invoice.ExpireParams{
// 		ID: payment.InvoiceID,
// 	}
// 	_, errs := pr.Xendit.Expire(&cancelInvoice)
// 	if errs != nil {
// 		return errs
// 	}
// 	return nil
// }
