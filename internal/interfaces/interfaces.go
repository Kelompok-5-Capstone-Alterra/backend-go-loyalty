package interfaces

import (
	"context"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type XenditConn interface {
	Create(data *invoice.CreateParams) (*xendit.Invoice, *xendit.Error)
	CreateWithContext(ctx context.Context, data *invoice.CreateParams) (*xendit.Invoice, *xendit.Error)
	Get(data *invoice.GetParams) (*xendit.Invoice, *xendit.Error)
	GetWithContext(ctx context.Context, data *invoice.GetParams) (*xendit.Invoice, *xendit.Error)
	Expire(data *invoice.ExpireParams) (*xendit.Invoice, *xendit.Error)
	ExpireWithContext(ctx context.Context, data *invoice.ExpireParams) (*xendit.Invoice, *xendit.Error)
	GetAll(data *invoice.GetAllParams) ([]xendit.Invoice, *xendit.Error)
	GetAllWithContext(ctx context.Context, data *invoice.GetAllParams) ([]xendit.Invoice, *xendit.Error)
}
