package entity

import (
	"testing"
	"time"

	pkg_dates "count_on_us/pkg/dates"

	"github.com/stretchr/testify/assert"
)

func TestAccountEntitySuccess(t *testing.T) {
	t.Run("should return a new account", func(t *testing.T) {
		parsedTime, _ := pkg_dates.StrToTime(time.Now().GoString())
		account, err := NewAccount(parsedTime, parsedTime, "this is a test", 10, CREDIT, OPENED, "1", 10)
		assert.Nil(t, err)
		assert.NotNil(t, account)
		assert.Equal(t, "this is a test", account.Description)
		assert.Equal(t, 10.0, account.Value)
		assert.Equal(t, OPENED, account.Status)
		assert.Equal(t, CREDIT, account.Type)
		assert.Equal(t, "1", account.AccountGroupID)
		assert.NotEmpty(t, account.ID)
	})
}

func TestNewAccounttEntityErrors(t *testing.T) {

	dueDate, _ := pkg_dates.StrToTime(time.Now().String())
	paymentDate, _ := pkg_dates.StrToTime(time.Now().String())
	description := "test"
	value := 10.0
	aType := CREDIT
	status := OPENED
	groupID := "1"
	installments := 1

	t.Run("should throw an error when account value is lesser or equal zero", func(t *testing.T) {
		account, err := NewAccount(dueDate, paymentDate, description, 0.00, aType, status, groupID, installments)
		assert.Nil(t, account)
		assert.EqualError(t, err, "account value cannot be equal or lesser than zero!")
	})

	t.Run("should throw an error when installments is lesser or equal than zero", func(t *testing.T) {
		account, err := NewAccount(dueDate, paymentDate, description, value, aType, status, groupID, 0)
		assert.Nil(t, account)
		assert.EqualError(t, err, "installments cannot be lesser or equal than zero!")
	})

	t.Run("should throw an error when description is empty", func(t *testing.T) {
		account, err := NewAccount(dueDate, paymentDate, "", value, aType, status, groupID, installments)
		assert.Nil(t, account)
		assert.EqualError(t, err, "description cannot be empty!")
	})

	t.Run("should throw an error when groupID is empty", func(t *testing.T) {
		account, err := NewAccount(dueDate, paymentDate, description, value, aType, status, "", installments)
		assert.Nil(t, account)
		assert.EqualError(t, err, "group ID cannot be empty!")
	})
}
