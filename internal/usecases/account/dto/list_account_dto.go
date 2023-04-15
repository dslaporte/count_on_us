package usecases

type ListInputDTO struct {
	Limit  int
	Offset int
}

type ListOutputDTO struct {
	ID           string
	DueDate      string
	PaymentDate  string
	Description  string
	Value        float64
	Type         string
	Status       string
	Installments int
}
