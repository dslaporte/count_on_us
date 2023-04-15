package usecases

import (
	entity "count_on_us/internal/entity/account"
	dto "count_on_us/internal/usecases/account/dto"
)

type ListAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewListUseCase(repository entity.AccountRepositoryInterface) *ListAccountUseCase {
	return &ListAccountUseCase{
		AccountRepository: repository,
	}
}

func (u *ListAccountUseCase) Execute(input dto.ListInputDTO) ([]dto.ListOutputDTO, error) {
	output := make([]dto.ListOutputDTO, 0)
	accounts, err := u.AccountRepository.List(input.Limit, input.Offset)
	if err != nil {
		return output, err
	}
	for _, account := range accounts {
		output = append(output, dto.ListOutputDTO{
			ID:           account.ID,
			DueDate:      account.DueDate,
			PaymentDate:  account.PaymentDate,
			Description:  account.Description,
			Value:        account.Value,
			Type:         string(account.Type),
			Status:       string(account.Status),
			Installments: account.Installments,
		})
	}
	return output, nil
}
