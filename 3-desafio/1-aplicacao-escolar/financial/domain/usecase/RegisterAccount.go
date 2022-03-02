package usecase

import "financial/domain/entity"

type RegisterAccount struct {
}

func NewRegisterAccount() *RegisterAccount {
	return &RegisterAccount{}
}

func (uc RegisterAccount) Execute(account *entity.Account) error {

	return nil
}
