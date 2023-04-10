package models

import (
	"database/sql"
)

type Account struct {
	ID           string        `db:"id"`
	Description  string        `db:"description"`
	DueDate      string        `db:"due_date"`
	PaymentDate  string        `db:"payment_date"`
	Value        float64       `db:"value"`
	Type         string        `db:"type"`
	Status       string        `db:"status"`
	OwnerID      string        `db:"owner_id"`
	GroupID      string        `db:"account_group_id"`
	Installments int           `db:"installments"`
	CreatedAt    *sql.NullTime `db:"created_at"`
	UpdatedAt    *sql.NullTime `db:"updated_at"` 
}
