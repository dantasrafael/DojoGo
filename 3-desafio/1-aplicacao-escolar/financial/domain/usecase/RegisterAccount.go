package usecase

type RegisterAccount struct {
}

func NewRegisterAccount() *RegisterAccount {
	return &RegisterAccount{}
}

func (uc RegisterAccount) Execute() error {
	return nil
}
