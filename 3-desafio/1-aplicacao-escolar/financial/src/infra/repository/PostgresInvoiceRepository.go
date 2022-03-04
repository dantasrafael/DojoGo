package repository

import (
	"context"
	"database/sql"
	"financial/src/domain/entity"
	"financial/src/domain/repository"
	"fmt"
)

const (
	insertInvoice = "INSERT INTO invoice (account_id, installment, due_date, value) VALUES ( $1, $2, $3, $4)"
)

type PostgresInvoiceRepository struct {
	db *sql.DB
}

func NewPostgresInvoiceRepository(db *sql.DB) repository.InvoiceRepository {
	return PostgresInvoiceRepository{db: db}
}

func (p PostgresInvoiceRepository) Save(ctx context.Context, invoice *entity.Invoice) error {
	row := p.db.QueryRowContext(ctx, insertInvoice, invoice.Account.ID, invoice.Installment, invoice.DueDate, invoice.Value)
	if row.Err() != nil {
		return fmt.Errorf("could not save invoice with account[%s]: %v", invoice.Account.ID, row.Err())
	}
	return nil
}
