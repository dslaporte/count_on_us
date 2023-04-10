package usecases

import entity "controle_conta/internal/entity/account"

type CreateAccountInputDTO struct {
	DueDate        string
	PaymentDate    string
	Description    string
	Value          float64
	AccountType    entity.AccountType
	Status         entity.AccountStatus
	AccountGroupID string
	Installments   int
}

type CreateAccoutOutputDTO struct {
	DueDate      string
	PaymentDate  string
	Description  string
	Value        float64
	Type         string
	Status       string
	Installments int
}
