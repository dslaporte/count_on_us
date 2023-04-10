package entity

type AccountRepositoryInterface interface {
	Save(account *Account) error
	List() ([]*Account, error)
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}
