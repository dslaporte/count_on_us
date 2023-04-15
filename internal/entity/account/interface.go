package entity

type AccountRepositoryInterface interface {
	Save(account *Account) error
	List(limit, offset int) ([]*Account, error)
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}
