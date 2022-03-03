package usecase

import (
	"context"
	"database/sql"
	"financial/domain/entity"
	"financial/domain/repository"
	"financial/infra/producers"
	repository2 "financial/infra/repository"
	"fmt"
	"time"
)

type RegisterAccount struct {
	accountRepository       repository.AccountRepository
	invoiceRepository       repository.InvoiceRepository
	registerAccountProducer producers.RegisterAccountProducer
}

func NewRegisterAccount(db *sql.DB) *RegisterAccount {
	return &RegisterAccount{
		accountRepository:       repository2.NewPostgresAccountRepository(db),
		invoiceRepository:       repository2.NewPostgresInvoiceRepository(db),
		registerAccountProducer: producers.NewRegisterAccountProducer(),
	}
}

func (uc RegisterAccount) Execute(ctx context.Context, account *entity.Account) error {
	account.Status = entity.ADIMPLENTE
	accountId, err := uc.accountRepository.Save(ctx, account)

	if err != nil {
		return fmt.Errorf("could not execute register account: %v", err)
	}
	account.ID = *accountId
	invoices := uc.splitInvoices(account)

	for _, invoice := range invoices {
		err := uc.invoiceRepository.Save(ctx, &invoice)
		if err != nil {
			return fmt.Errorf("could not save invoice[%d]: %v", invoice.Installment, err)
		}
	}

	go uc.registerAccountProducer.Send(ctx, account)
	return nil
}

func (uc RegisterAccount) splitInvoices(account *entity.Account) []entity.Invoice {
	invoiceValue := account.Total / float64(account.Installments)
	invoices := make([]entity.Invoice, 0, account.Installments)
	for idx := 0; idx < int(account.Installments); idx++ {
		monthToAdd := idx + 1
		invoice := entity.Invoice{
			Account:     account,
			Installment: uint8(monthToAdd),
			CreatedAt:   time.Now(),
			DueDate:     time.Now().AddDate(0, monthToAdd, 0),
			Value:       invoiceValue,
		}
		invoices = append(invoices, invoice)
	}
	return invoices
}
