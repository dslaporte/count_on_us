package entity

import (
	commons "controle_conta/internal/commons/strings"
	"errors"

	"github.com/google/uuid"
)

type AccountType string

type OutgoingType string

type AccountStatus string

const (
	FIX      OutgoingType = "fixed"
	VARIABLE OutgoingType = "variable"
)

const (
	CREDIT AccountType = "credit" //credito
	DEBIT  AccountType = "debit"  //debito
)

const (
	PAID      AccountStatus = "paid"      //pago
	OPENED    AccountStatus = "opened"    //em aberto
	OVERDUE   AccountStatus = "overdue"   //em atraso
	SCHEDULED AccountStatus = "scheduled" //agendada
)

type Account struct {
	ID             string
	DueDate        string //vencimento
	PaymentDate    string //pagamento
	Description    string
	Value          float64
	Type           AccountType
	Status         AccountStatus
	AccountGroupID string
	OwnerID        string
	Installments   int
}

func NewAccount(
	dueDate string,
	paymentDate string,
	description string,
	value float64,
	accountType AccountType,
	status AccountStatus,
	groupID string,
	installments int) (*Account, error) {
	a := &Account{
		ID:             uuid.NewString(),
		DueDate:        dueDate,
		PaymentDate:    paymentDate,
		Description:    description,
		Value:          value,
		Type:           accountType,
		Status:         status,
		AccountGroupID: groupID,
		Installments:   installments,
	}
	if err := a.IsValid(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Account) IsValid() error {
	if a.Value <= 0 {
		return errors.New("account value cannot be equal or lesser than zero!")
	}
	if a.Installments <= 0 {
		return errors.New("installments cannot be lesser or equal than zero!")
	}
	if commons.IsEmpty(a.AccountGroupID) {
		return errors.New("group ID cannot be empty!")
	}
	if commons.IsEmpty(a.Description) {
		return errors.New("description cannot be empty!")
	}
	return nil
}

// LANÇAMENTO:

// DATA | 100,00  | PARTICULAR | DANIEL  | DESPESA | COMPROVANTE
// DATA | 250,00  | PARTICULAR | NATALIA | CREDITO |
// DATA | 1200,00 | CARPE DIEN | TODOS   | DESPESA |

// CATEGODate
//   - ID
//   - DESCRIÇÃO
//      - LAZER
// 	 - CULTURA
// 	 - SAUDE

// SUB-CATEGORIA:
//    - ID
//    - DESCRIÇAOd
//    - ID CATEGORIA (SAUDE)
//      - MEDICAMENTO
// 	 - PLANO DE SAUDE
// 	 - MEDICINA ALTERNATIVA

// TIPO DE DESPESA: FIXA | VARIAVEL
