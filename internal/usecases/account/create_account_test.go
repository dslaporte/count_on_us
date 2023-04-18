package usecases

import (
	"fmt"
	"testing"
	"time"

	entity "count_on_us/internal/entity/account"
	usecases "count_on_us/internal/usecases/account/dto"
	pkg_dates "count_on_us/pkg/dates"
	repository "count_on_us/tests/mocks/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	account, err := entity.NewAccount(time.Now(), time.Now(), "teste", 100, entity.DEBIT, entity.OPENED, "1", 1)
	assert.Nil(t, err)

	t.Run("should return errors when try to create a new Account entitty", func(t *testing.T) {

		type accountErrorTestcases struct {
			description   string
			input         usecases.CreateAccountInputDTO
			expectedError string
		}

		accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
		tests := []accountErrorTestcases{
			{
				description: "should return an error when try to create a entity with value equal or lesser than zero",
				input: usecases.CreateAccountInputDTO{
					DueDate:        account.DueDate.String(),
					PaymentDate:    account.PaymentDate.String(),
					Description:    "test",
					Value:          -100,
					AccountType:    entity.CREDIT,
					Status:         entity.OPENED,
					AccountGroupID: "1",
					Installments:   1,
				},
				expectedError: "account value cannot be equal or lesser than zero!",
			},
			{
				description: "should return an error when installments are lesser or equal than zero",
				input: usecases.CreateAccountInputDTO{
					DueDate:        account.DueDate.String(),
					PaymentDate:    account.PaymentDate.String(),
					Description:    "test",
					Value:          100,
					AccountType:    entity.CREDIT,
					Status:         entity.OPENED,
					AccountGroupID: "1",
					Installments:   0,
				},
				expectedError: "installments cannot be lesser or equal than zero!",
			},
			{
				description: "should return an error when groupd ID is empty",
				input: usecases.CreateAccountInputDTO{
					DueDate:        account.DueDate.String(),
					PaymentDate:    account.PaymentDate.String(),
					Description:    "test",
					Value:          100,
					AccountType:    entity.CREDIT,
					Status:         entity.OPENED,
					AccountGroupID: "",
					Installments:   1,
				},
				expectedError: "group ID cannot be empty!",
			}, {
				description: "should return an error when description is empty",
				input: usecases.CreateAccountInputDTO{
					DueDate:        account.DueDate.String(),
					PaymentDate:    account.PaymentDate.String(),
					Description:    "",
					Value:          100,
					AccountType:    entity.CREDIT,
					Status:         entity.OPENED,
					AccountGroupID: "1",
					Installments:   1,
				},
				expectedError: "description cannot be empty!",
			},
		}

		for _, test := range tests {
			createAccountUseCase, err := NewCreateAccountUseCase(accountRepository).Execute(test.input)
			assert.Nil(t, createAccountUseCase)
			assert.NotNil(t, err)
			assert.EqualError(t, err, test.expectedError)
		}
	})

	t.Run("should return a repository error", func(t *testing.T) {
		expectedErrorMsg := "cannot save record"
		accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
		accountRepository.EXPECT().Save(gomock.Any()).Return(fmt.Errorf(expectedErrorMsg))
		createAccountUseCase, err := NewCreateAccountUseCase(accountRepository).Execute(usecases.CreateAccountInputDTO{
			DueDate:        account.DueDate.String(),
			PaymentDate:    account.PaymentDate.String(),
			Description:    account.Description,
			Value:          account.Value,
			AccountType:    account.Type,
			Status:         account.Status,
			AccountGroupID: account.AccountGroupID,
			Installments:   account.Installments,
		})
		assert.Nil(t, createAccountUseCase)
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErrorMsg)
	})

	t.Run("should not return an error to create a new account", func(t *testing.T) {
		accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
		accountRepository.EXPECT().Save(gomock.Any()).Return(nil)

		accountInputDTO := usecases.CreateAccountInputDTO{
			DueDate:        account.DueDate.String(),
			PaymentDate:    account.PaymentDate.String(),
			Description:    account.Description,
			Value:          account.Value,
			AccountType:    account.Type,
			Status:         account.Status,
			AccountGroupID: account.AccountGroupID,
			Installments:   account.Installments,
		}

		dueDate, _ := pkg_dates.StrToTime(accountInputDTO.DueDate)
		paymentDate, _ := pkg_dates.StrToTime(accountInputDTO.PaymentDate)
		accountOutputDTO, err := NewCreateAccountUseCase(accountRepository).Execute(accountInputDTO)
		assert.Nil(t, err)
		assert.NotNil(t, accountOutputDTO)
		assert.NotEmpty(t, accountOutputDTO.ID)
		assert.Equal(t, accountOutputDTO.Description, accountInputDTO.Description)
		assert.Equal(t, accountOutputDTO.DueDate, dueDate)
		assert.Equal(t, accountOutputDTO.PaymentDate, paymentDate)
		assert.Equal(t, accountOutputDTO.Value, accountInputDTO.Value)
		assert.Equal(t, accountOutputDTO.Type, string(accountInputDTO.AccountType))
		assert.Equal(t, accountOutputDTO.Status, string(accountInputDTO.Status))
		assert.Equal(t, accountOutputDTO.Installments, accountInputDTO.Installments)
	})
}
