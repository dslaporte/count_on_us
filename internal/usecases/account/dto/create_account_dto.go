package usecases

import (
	entity "count_on_us/internal/entity/account"
	"time"
)

type CreateAccountInputDTO struct {
	DueDate        string               `json:"due_date" validate:"required,rfc3339"`
	PaymentDate    string               `json:"payment_date" validate:"required,rfc3339"`
	Description    string               `json:"description" validate:"required"`
	Value          float64              `json:"value" validate:"gt=0,required"`
	AccountType    entity.AccountType   `json:"account_type" validate:"required"`
	Status         entity.AccountStatus `json:"status" validate:"required,enum"`
	AccountGroupID string               `json:"group_id" validate:"required"`
	Installments   int                  `json:"installments" validate:"gt=0,required"`
}

type CreateAccoutOutputDTO struct {
	ID           string    `json:"id"`
	DueDate      time.Time `json:"due_date"`
	PaymentDate  time.Time `json:"payment_date"`
	Description  string    `json:"description"`
	Value        float64   `json:"value"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	Installments int       `json:"installments"`
}
