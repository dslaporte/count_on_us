package usecases

import (
	entity "count_on_us/src/internal/entity/account"
)

type CreateAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewCreateAccountUseCase(
	repository entity.AccountRepositoryInterface,
) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository: repository,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccoutOutputDTO, error) {
	account, err := entity.NewAccount(
		input.DueDate,
		input.PaymentDate,
		input.Description,
		input.Value,
		input.AccountType,
		input.Status,
		input.AccountGroupID,
		input.Installments,
	)
	if err != nil {
		return nil, err
	}
	if err = uc.AccountRepository.Save(account); err != nil {
		return nil, err
	}
	return &CreateAccoutOutputDTO{
		DueDate:      account.DueDate,
		PaymentDate:  account.PaymentDate,
		Description:  account.Description,
		Value:        account.Value,
		Type:         string(account.Type),
		Status:       string(account.Status),
		Installments: input.Installments,
	}, nil
}
