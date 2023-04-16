package usecases

import (
	entity "count_on_us/internal/entity/account"

	dto "count_on_us/internal/usecases/account/dto"
	pkg_dates "count_on_us/pkg/dates"
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

func (uc *CreateAccountUseCase) Execute(input dto.CreateAccountInputDTO) (*dto.CreateAccoutOutputDTO, error) {
	parsedDueDate, _ := pkg_dates.StrToTime(input.DueDate)
	parsedPaymentDate, _ := pkg_dates.StrToTime(input.PaymentDate)
	account, err := entity.NewAccount(
		parsedDueDate,
		parsedPaymentDate,
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
	return &dto.CreateAccoutOutputDTO{
		ID:           account.ID,
		DueDate:      account.DueDate,
		PaymentDate:  account.PaymentDate,
		Description:  account.Description,
		Value:        account.Value,
		Type:         string(account.Type),
		Status:       string(account.Status),
		Installments: input.Installments,
	}, nil
}
