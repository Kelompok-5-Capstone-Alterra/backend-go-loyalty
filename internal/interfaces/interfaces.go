package interfaces

import (
	"context"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
	"github.com/xendit/xendit-go/invoice"
)

type XenditInvoice interface {
	Create(data *invoice.CreateParams) (*xendit.Invoice, *xendit.Error)
	CreateWithContext(ctx context.Context, data *invoice.CreateParams) (*xendit.Invoice, *xendit.Error)
	Get(data *invoice.GetParams) (*xendit.Invoice, *xendit.Error)
	GetWithContext(ctx context.Context, data *invoice.GetParams) (*xendit.Invoice, *xendit.Error)
	Expire(data *invoice.ExpireParams) (*xendit.Invoice, *xendit.Error)
	ExpireWithContext(ctx context.Context, data *invoice.ExpireParams) (*xendit.Invoice, *xendit.Error)
	GetAll(data *invoice.GetAllParams) ([]xendit.Invoice, *xendit.Error)
	GetAllWithContext(ctx context.Context, data *invoice.GetAllParams) ([]xendit.Invoice, *xendit.Error)
}

type XenditEWallet interface {
	CreatePayment(data *ewallet.CreatePaymentParams) (*xendit.EWallet, *xendit.Error)
	CreatePaymentWithContext(ctx context.Context, data *ewallet.CreatePaymentParams) (*xendit.EWallet, *xendit.Error)
	GetPaymentStatus(data *ewallet.GetPaymentStatusParams) (*xendit.EWallet, *xendit.Error)
	GetPaymentStatusWithContext(ctx context.Context, data *ewallet.GetPaymentStatusParams) (*xendit.EWallet, *xendit.Error)
	CreateEWalletCharge(data *ewallet.CreateEWalletChargeParams) (*xendit.EWalletCharge, *xendit.Error)
	CreateEWalletChargeWithContext(ctx context.Context, data *ewallet.CreateEWalletChargeParams) (*xendit.EWalletCharge, *xendit.Error)
	GetEWalletChargeStatus(data *ewallet.GetEWalletChargeStatusParams) (*xendit.EWalletCharge, *xendit.Error)
	GetEWalletChargeStatusWithContext(ctx context.Context, data *ewallet.GetEWalletChargeStatusParams) (*xendit.EWalletCharge, *xendit.Error)
}
