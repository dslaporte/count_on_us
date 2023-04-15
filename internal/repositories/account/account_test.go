package repositories

import (
	entity "count_on_us/internal/entity/account"
	"count_on_us/tests"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	db := tests.SetupDB()
	defer db.Close()

	account, err := entity.NewAccount(
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		"account for test",
		100.00,
		entity.DEBIT,
		entity.OPENED,
		"1",
		1,
	)
	assert.Nil(t, err)

	repository := NewAccountRepository(db)
	assert.NotNil(t, repository)
	err = repository.Save(account)
	assert.Nil(t, err)

	accountFound, err := repository.FindByID(account.ID)
	assert.Nil(t, err)
	assert.Equal(t, account.Description, accountFound.Description)
	assert.Equal(t, account.DueDate, accountFound.DueDate)
	assert.Equal(t, account.ID, accountFound.ID)
	assert.Equal(t, account.Type, accountFound.Type)
	assert.Equal(t, account.Installments, accountFound.Installments)
	assert.Equal(t, account.Value, accountFound.Value)
	assert.Equal(t, account.PaymentDate, accountFound.PaymentDate)
}

func TestUpdateAccount(t *testing.T) {
	db := tests.SetupDB()
	defer db.Close()

	account, err := entity.NewAccount(
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		"account for test",
		100.00,
		entity.DEBIT,
		entity.OPENED,
		"1",
		1,
	)
	assert.Nil(t, err)

	repository := NewAccountRepository(db)
	assert.NotNil(t, repository)
	err = repository.Save(account)
	assert.Nil(t, err)

	expectedDescription := "blabla"
	expectedValue := 500.00

	account.Description = expectedDescription
	account.Value = expectedValue

	err = repository.Update(account)
	assert.Nil(t, err)

	accountFound, err := repository.FindByID(account.ID)
	assert.Nil(t, err)
	assert.Equal(t, account.Description, accountFound.Description)
	assert.Equal(t, account.DueDate, accountFound.DueDate)
	assert.Equal(t, account.ID, accountFound.ID)
	assert.Equal(t, account.Type, accountFound.Type)
	assert.Equal(t, account.Installments, accountFound.Installments)
	assert.Equal(t, account.Value, accountFound.Value)
	assert.Equal(t, account.PaymentDate, accountFound.PaymentDate)

	assert.Equal(t, accountFound.Value, expectedValue)
	assert.Equal(t, accountFound.Description, expectedDescription)
}

func TestListAccounts(t *testing.T) {
	db := tests.SetupDB()
	defer db.Close()

	repository := NewAccountRepository(db)
	assert.NotNil(t, repository)

	accountsCreated := make([]entity.Account, 0)
	for i := 1; i < 25; i++ {
		account, err := entity.NewAccount(
			time.Now().Format(time.RFC3339),
			time.Now().Format(time.RFC3339),
			fmt.Sprintf("Account %d", i),
			rand.Float64(),
			entity.CREDIT,
			entity.OPENED,
			"1",
			1,
		)
		assert.Nil(t, err)
		accountsCreated = append(accountsCreated, *account)
		err = repository.Save(account)
		assert.Nil(t, err)
	}

	//first page
	accounts, err := repository.List(10, 1)
	assert.Nil(t, err)
	assert.Len(t, accounts, 10)
	for idx, account := range accounts {
		assert.Equal(t, &accountsCreated[idx], account)
	}

	//second page
	accounts, err = repository.List(10, 2)
	assert.Nil(t, err)
	assert.Len(t, accounts, 10)
	for idx, account := range accounts {
		assert.Equal(t, &accountsCreated[10+idx], account)
	}

	//third page
	accounts, err = repository.List(10, 3)
	assert.Nil(t, err)
	assert.Len(t, accounts, 4)
	for idx, account := range accounts {
		assert.Equal(t, &accountsCreated[20+idx], account)
	}
}
