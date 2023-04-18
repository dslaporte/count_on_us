package usecases

import (
	entity "count_on_us/internal/entity/account"
	usecases "count_on_us/internal/usecases/account/dto"
	repository "count_on_us/tests/mocks/repository"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAccountUseCase(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	accountRepository := repository.NewMockAccountRepositoryInterface(ctrl)
	assert.NotNil(t, accountRepository)
	account := entity.Account{
		ID:             uuid.NewString(),
		DueDate:        time.Now(),
		PaymentDate:    time.Now(),
		Description:    "account test",
		Value:          200,
		Type:           entity.CREDIT,
		Status:         entity.OPENED,
		AccountGroupID: "1",
		OwnerID:        "",
		Installments:   0,
	}
	t.Run("should return an error when try to update an account", func(t *testing.T) {
		expectedError := "database error"
		inputDTO := usecases.UpdateAccountInputDTO{
			DueDate:        account.DueDate,
			PaymentDate:    account.PaymentDate,
			Description:    account.Description,
			Value:          account.Value,
			AccountType:    account.Type,
			Status:         account.Status,
			AccountGroupID: account.AccountGroupID,
		}
		accountUpdateUsecase := NewUpdateAccountUseCase(accountRepository)
		accountRepository.EXPECT().Update(&account).Return(fmt.Errorf(expectedError))
		accountUpdateOutputDTO, err := accountUpdateUsecase.Execute(account.ID, inputDTO)
		assert.NotNil(t, err)
		assert.Nil(t, accountUpdateOutputDTO)
		assert.EqualError(t, err, expectedError)
	})

	t.Run("should return an updated record", func(t *testing.T) {
		inputDTO := usecases.UpdateAccountInputDTO{
			DueDate:        account.DueDate,
			PaymentDate:    account.PaymentDate,
			Description:    account.Description,
			Value:          account.Value,
			AccountType:    account.Type,
			Status:         account.Status,
			AccountGroupID: account.AccountGroupID,
		}

		expectedUpdateOutputDTO := usecases.UpdateAccountOutputDTO{
			ID:          account.ID,
			DueDate:     account.DueDate,
			PaymentDate: account.PaymentDate,
			Description: account.Description,
			Value:       account.Value,
			AccountType: account.Type,
			Status:      account.Status,
			Group: &usecases.UpdateGroupOutputDTO{
				ID:          account.AccountGroupID,
				Description: "",
			},
		}

		accountUpdateUsecase := NewUpdateAccountUseCase(accountRepository)
		accountRepository.EXPECT().Update(&account).Return(nil).AnyTimes()
		accountRepository.EXPECT().FindByID(account.ID).Return(&account, nil).AnyTimes()
		accountUpdateOutputDTO, err := accountUpdateUsecase.Execute(account.ID, inputDTO)
		assert.Nil(t, err)
		assert.NotNil(t, accountUpdateOutputDTO)
		assert.Equal(t, &expectedUpdateOutputDTO, accountUpdateOutputDTO)
	})
}
