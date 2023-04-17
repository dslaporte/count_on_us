package usecases

import (
	entity "count_on_us/internal/entity/account"
	dto "count_on_us/internal/usecases/account/dto"
)

type FindAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewFindAccountUseCase(repository entity.AccountRepositoryInterface) *FindAccountUseCase {
	return &FindAccountUseCase{
		AccountRepository: repository,
	}
}

func (u *FindAccountUseCase) Execute(id string) (*dto.FindAccoutOutputDTO, error) {
	account, err := u.AccountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.FindAccoutOutputDTO{
		ID:           account.ID,
		DueDate:      account.DueDate,
		PaymentDate:  account.PaymentDate,
		Description:  account.Description,
		Value:        account.Value,
		Type:         string(account.Type),
		Status:       string(account.Status),
		Installments: account.Installments,
	}, nil
}
