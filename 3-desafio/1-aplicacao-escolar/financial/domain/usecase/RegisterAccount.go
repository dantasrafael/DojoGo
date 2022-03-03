package usecase

import (
	"context"
	"financial/domain/entity"
)

type RegisterAccount struct {
}

func NewRegisterAccount() *RegisterAccount {
	return &RegisterAccount{}
}

func (uc RegisterAccount) Execute(ctx context.Context, account *entity.Account) error {

	return nil
}
