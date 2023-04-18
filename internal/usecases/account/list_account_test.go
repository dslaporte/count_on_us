package usecases

import (
	entity "count_on_us/internal/entity/account"
	usecases "count_on_us/internal/usecases/account/dto"
	repository "count_on_us/tests/mocks/repository"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListAccountUseCase(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
	t.Run("should return an error when list accounts", func(t *testing.T) {
		expectedError := "database error"
		expectedAccountResult := make([]usecases.ListOutputDTO, 0)
		accountListUsecase := NewListUseCase(accountRepository)
		accountRepository.EXPECT().List(1, 1).Return(make([]*entity.Account, 0), fmt.Errorf(expectedError))
		accountListOutputDTO, err := accountListUsecase.Execute(usecases.ListInputDTO{Limit: 1, Offset: 1})
		assert.NotNil(t, err)
		assert.NotNil(t, accountListOutputDTO)
		assert.Equal(t, expectedAccountResult, accountListOutputDTO)
	})

	t.Run("should not return a error when records are found", func(t *testing.T) {
		account1, err := entity.NewAccount(time.Now(), time.Now(), "account 1", 100, entity.CREDIT, entity.OPENED, "1", 1)
		assert.Nil(t, err)
		account2, err := entity.NewAccount(time.Now(), time.Now(), "account 2", 200, entity.DEBIT, entity.OPENED, "1", 1)
		assert.Nil(t, err)

		expectedAccountResult := []usecases.ListOutputDTO{
			{
				ID:           account1.ID,
				DueDate:      account1.DueDate,
				PaymentDate:  account1.PaymentDate,
				Description:  account1.Description,
				Value:        account1.Value,
				Type:         string(account1.Type),
				Status:       string(account1.Status),
				Installments: account1.Installments,
			},
			{
				ID:           account2.ID,
				DueDate:      account2.DueDate,
				PaymentDate:  account2.PaymentDate,
				Description:  account2.Description,
				Value:        account2.Value,
				Type:         string(account2.Type),
				Status:       string(account2.Status),
				Installments: account2.Installments,
			},
		}

		accountListUsecase := NewListUseCase(accountRepository)
		accountRepository.EXPECT().List(2, 1).Return([]*entity.Account{account1, account2}, nil)
		accountListOutputDTO, err := accountListUsecase.Execute(usecases.ListInputDTO{Limit: 2, Offset: 1})
		assert.Nil(t, err)
		assert.NotNil(t, accountListOutputDTO)
		assert.Equal(t, expectedAccountResult, accountListOutputDTO)
	})
}
