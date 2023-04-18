package usecases

import (
	"fmt"
	"testing"
	"time"

	entity "count_on_us/internal/entity/account"
	usecases "count_on_us/internal/usecases/account/dto"
	repository "count_on_us/tests/mocks/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindAccountUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account, err := entity.NewAccount(time.Now(), time.Now(), "teste", 100, entity.DEBIT, entity.OPENED, "1", 1)
	assert.Nil(t, err)
	assert.NotNil(t, account)

	t.Run("should return an error when record was not found", func(t *testing.T) {
		expectedErrorMsg := "record not found"
		accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
		accountRepository.EXPECT().FindByID(account.ID).Return(nil, fmt.Errorf(expectedErrorMsg))
		findAccoutOutputDTO, err := NewFindAccountUseCase(accountRepository).Execute(account.ID)
		assert.Nil(t, findAccoutOutputDTO)
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErrorMsg)
	})

	t.Run("should return an account", func(t *testing.T) {
		accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
		accountRepository.EXPECT().FindByID(account.ID).Return(account, nil)

		findAccountOutputDTO := usecases.FindAccountOutputDTO{
			ID:           account.ID,
			DueDate:      account.DueDate,
			PaymentDate:  account.PaymentDate,
			Description:  account.Description,
			Value:        account.Value,
			Type:         string(account.Type),
			Status:       string(account.Status),
			Installments: account.Installments,
		}

		accountOutputDTO, err := NewFindAccountUseCase(accountRepository).Execute(account.ID)
		assert.Nil(t, err)
		assert.NotNil(t, accountOutputDTO)
		assert.NotEmpty(t, accountOutputDTO.ID)
		assert.Equal(t, accountOutputDTO.Description, findAccountOutputDTO.Description)
		assert.Equal(t, accountOutputDTO.DueDate, findAccountOutputDTO.DueDate)
		assert.Equal(t, accountOutputDTO.PaymentDate, findAccountOutputDTO.PaymentDate)
		assert.Equal(t, accountOutputDTO.Value, findAccountOutputDTO.Value)
		assert.Equal(t, accountOutputDTO.Type, findAccountOutputDTO.Type)
		assert.Equal(t, accountOutputDTO.Status, findAccountOutputDTO.Status)
		assert.Equal(t, accountOutputDTO.Installments, findAccountOutputDTO.Installments)
	})
}
