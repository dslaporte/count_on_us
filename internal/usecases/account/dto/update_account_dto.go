package usecases

import (
	entity "count_on_us/internal/entity/account"
	"time"
)

type UpdateAccountInputDTO struct {
	DueDate        time.Time            `json:"due_date" validate:"required,rfc3339"`
	PaymentDate    time.Time            `json:"payment_date" validate:"required,rfc3339"`
	Description    string               `json:"description" validate:"required"`
	Value          float64              `json:"value" validate:"gt=0,required"`
	AccountType    entity.AccountType   `json:"account_type" validate:"required"`
	Status         entity.AccountStatus `json:"status" validate:"required,enum"`
	AccountGroupID string               `json:"group_id" validate:"required"`
}

type UpdateGroupOutputDTO struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateAccountOutputDTO struct {
	ID          string                `json:"id"`
	DueDate     time.Time             `json:"due_date"`
	PaymentDate time.Time             `json:"payment_date"`
	Description string                `json:"description"`
	Value       float64               `json:"value"`
	AccountType entity.AccountType    `json:"account_type"`
	Status      entity.AccountStatus  `json:"status"`
	Group       *UpdateGroupOutputDTO `json:"group,omitempty"`
}
