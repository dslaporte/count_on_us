package usecases

import "time"

type FindAccoutOutputDTO struct {
	ID           string    `json:"id"`
	DueDate      time.Time `json:"due_date"`
	PaymentDate  time.Time `json:"payment_date"`
	Description  string    `json:"description"`
	Value        float64   `json:"value"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	Installments int       `json:"installments"`
}
