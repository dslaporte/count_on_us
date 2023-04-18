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

	t.Run("should insert a new account", func(t *testing.T) {
		account, err := entity.NewAccount(
			time.Now(),
			time.Now(),
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
		assert.Equal(t, account.ID, accountFound.ID)
		assert.Equal(t, account.Type, accountFound.Type)
		assert.Equal(t, account.Installments, accountFound.Installments)
		assert.Equal(t, account.Value, accountFound.Value)
	})
}

func TestUpdateAccount(t *testing.T) {
	db := tests.SetupDB()
	defer db.Close()

	account, err := entity.NewAccount(
		time.Now(),
		time.Now(),
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
	assert.Equal(t, account.ID, accountFound.ID)
	assert.Equal(t, account.Type, accountFound.Type)
	assert.Equal(t, account.Installments, accountFound.Installments)
	assert.Equal(t, account.Value, accountFound.Value)

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
			time.Now(),
			time.Now(),
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
		assert.Equal(t, accountsCreated[idx].ID, account.ID)
		assert.Equal(t, accountsCreated[idx].AccountGroupID, account.AccountGroupID)
		assert.Equal(t, accountsCreated[idx].Description, account.Description)
		assert.Equal(t, accountsCreated[idx].Installments, account.Installments)
		assert.Equal(t, accountsCreated[idx].OwnerID, account.OwnerID)
		assert.Equal(t, accountsCreated[idx].Status, account.Status)
		assert.Equal(t, accountsCreated[idx].Type, account.Type)
		assert.Equal(t, accountsCreated[idx].Value, account.Value)
		// assert.Equal(t, &accountsCreated[idx], account)
	}

	//second page
	accounts, err = repository.List(10, 2)
	assert.Nil(t, err)
	assert.Len(t, accounts, 10)
	for idx, account := range accounts {
		assert.Equal(t, accountsCreated[10+idx].ID, account.ID)
		assert.Equal(t, accountsCreated[10+idx].AccountGroupID, account.AccountGroupID)
		assert.Equal(t, accountsCreated[10+idx].Description, account.Description)
		assert.Equal(t, accountsCreated[10+idx].Installments, account.Installments)
		assert.Equal(t, accountsCreated[10+idx].OwnerID, account.OwnerID)
		assert.Equal(t, accountsCreated[10+idx].Status, account.Status)
		assert.Equal(t, accountsCreated[10+idx].Type, account.Type)
		assert.Equal(t, accountsCreated[10+idx].Value, account.Value)
	}

	//third page
	accounts, err = repository.List(10, 3)
	assert.Nil(t, err)
	assert.Len(t, accounts, 4)
	for idx, account := range accounts {
		assert.Equal(t, accountsCreated[20+idx].ID, account.ID)
		assert.Equal(t, accountsCreated[20+idx].AccountGroupID, account.AccountGroupID)
		assert.Equal(t, accountsCreated[20+idx].Description, account.Description)
		assert.Equal(t, accountsCreated[20+idx].Installments, account.Installments)
		assert.Equal(t, accountsCreated[20+idx].OwnerID, account.OwnerID)
		assert.Equal(t, accountsCreated[20+idx].Status, account.Status)
		assert.Equal(t, accountsCreated[20+idx].Type, account.Type)
		assert.Equal(t, accountsCreated[20+idx].Value, account.Value)
	}
}
