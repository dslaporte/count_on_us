package usecases

import (
	entity "count_on_us/internal/entity/account"
	dto "count_on_us/internal/usecases/account/dto"
)

type UpdateAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewUpdateAccountUseCase(repository entity.AccountRepositoryInterface) *UpdateAccountUseCase {
	return &UpdateAccountUseCase{
		AccountRepository: repository,
	}
}

func (u *UpdateAccountUseCase) Execute(id string, inputDTO dto.UpdateAccountInputDTO) (*dto.UpdateAccountOutputDTO, error) {
	err := u.AccountRepository.Update(&entity.Account{
		ID:             id,
		DueDate:        inputDTO.DueDate,
		PaymentDate:    inputDTO.PaymentDate,
		Description:    inputDTO.Description,
		Value:          inputDTO.Value,
		Type:           inputDTO.AccountType,
		Status:         inputDTO.Status,
		AccountGroupID: inputDTO.AccountGroupID,
	})
	if err != nil {
		return nil, err
	}
	account, err := u.AccountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateAccountOutputDTO{
		ID:          account.ID,
		DueDate:     account.DueDate,
		PaymentDate: account.PaymentDate,
		Description: account.Description,
		Value:       account.Value,
		AccountType: account.Type,
		Status:      account.Status,
		Group: &dto.UpdateGroupOutputDTO{
			ID:          account.AccountGroupID,
			Description: "",
		},
	}, nil
}
