package usecases

import "time"

type ListInputDTO struct {
	Limit  int
	Offset int
}

type ListOutputDTO struct {
	ID           string
	DueDate      time.Time
	PaymentDate  time.Time
	Description  string
	Value        float64
	Type         string
	Status       string
	Installments int
}
