package paymentRepository

import (
	"backend-go-loyalty/internal/entity"
	"backend-go-loyalty/internal/interfaces"
	"backend-go-loyalty/pkg/utils"
	"context"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	CreateInvoice(ctx context.Context, req entity.Transaction, user entity.User) (*xendit.Invoice, error)
	InsertInvoiceData(ctx context.Context, req *xendit.Invoice, transactionID uint64) error
}

type paymentRepository struct {
	Xendit interfaces.XenditConn
	db     *gorm.DB
}

func NewPaymentRepository(xen interfaces.XenditConn, db *gorm.DB) paymentRepository {
	return paymentRepository{
		Xendit: xen,
		db:     db,
	}
}

func (pr paymentRepository) CreateInvoice(ctx context.Context, req entity.Transaction, user entity.User) (*xendit.Invoice, error) {
	customer := xendit.InvoiceCustomer{
		GivenNames:   user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Address:      nil,
	}
	items := []xendit.InvoiceItem{
		{
			Name:     req.Product.Name,
			Price:    float64(req.Product.Price),
			Quantity: 1,
		},
	}
	fee := xendit.InvoiceFee{
		Type:  "ADMIN",
		Value: 5000,
	}
	NotificationType := []string{"whatsapp", "email"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceReminder: NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceExpired:  NotificationType,
	}

	invoiceParams := invoice.CreateParams{
		ExternalID:                     utils.CreateExternalID(req.ID),
		Amount:                         float64(req.Amount + 5000),
		Description:                    "Invoice Pemesanan",
		InvoiceDuration:                3600,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		SuccessRedirectURL:             "https://www.google.com",
		FailureRedirectURL:             "https://www.google.com",
		Currency:                       "IDR",
		Items:                          items,
		Fees:                           []xendit.InvoiceFee{fee},
	}
	invo, err := pr.Xendit.Create(&invoiceParams)
	return invo, err
}
func (pr paymentRepository) InsertInvoiceData(ctx context.Context, req *xendit.Invoice, transactionID uint64) error {
	invoice := entity.PaymentInvoice{
		TransactionID: transactionID,
		URL:           req.InvoiceURL,
		Amount:        req.Amount,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := pr.db.Create(&invoice).Error
	return err
}
