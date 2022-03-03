package repository

import (
	"context"
	"financial/domain/entity"
)

type InvoiceRepository interface {
	Save(ctx context.Context, invoice *entity.Invoice) error
}
