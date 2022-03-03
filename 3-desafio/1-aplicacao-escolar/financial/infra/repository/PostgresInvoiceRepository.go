package repository

import (
	"context"
	"database/sql"
	"financial/domain/entity"
	"financial/domain/repository"
	"fmt"
)

const (
	insertInvoice = "INSERT INTO account (account_id, installment, due_date, value) VALUES ( $1, $2, $3, $4)"
)

type PostgresInvoiceRepository struct {
	db *sql.DB
}

func NewPostgresInvoiceRepository(db *sql.DB) repository.InvoiceRepository {
	return PostgresInvoiceRepository{db: db}
}

func (p PostgresInvoiceRepository) Save(ctx context.Context, invoice *entity.Invoice) error {
	err := p.db.QueryRowContext(ctx, insertInvoice, invoice.Account.ID, invoice.Installment, invoice.DueDate, invoice.Value)
	if err != nil {
		return fmt.Errorf("could not save invoice with account[%s]: %v", invoice.Account.ID, err)
	}
	return nil
}
