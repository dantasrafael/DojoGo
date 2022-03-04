package usecase

import (
	"context"
	"database/sql"
	entity2 "financial/src/domain/entity"
	repository2 "financial/src/domain/repository"
	"financial/src/infra/producers"
	repository3 "financial/src/infra/repository"
	"fmt"
	"time"
)

type RegisterAccount struct {
	accountRepository       repository2.AccountRepository
	invoiceRepository       repository2.InvoiceRepository
	registerAccountProducer producers.RegisterAccountProducer
}

func NewRegisterAccount(db *sql.DB) *RegisterAccount {
	return &RegisterAccount{
		accountRepository:       repository3.NewPostgresAccountRepository(db),
		invoiceRepository:       repository3.NewPostgresInvoiceRepository(db),
		registerAccountProducer: producers.NewRegisterAccountProducer(),
	}
}

func (uc RegisterAccount) Execute(ctx context.Context, account *entity2.Account) error {
	account.Status = entity2.ADIMPLENTE
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

func (uc RegisterAccount) splitInvoices(account *entity2.Account) []entity2.Invoice {
	invoiceValue := account.Total / float64(account.Installments)
	invoices := make([]entity2.Invoice, 0, account.Installments)
	for idx := 0; idx < int(account.Installments); idx++ {
		monthToAdd := idx + 1
		invoice := entity2.Invoice{
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
