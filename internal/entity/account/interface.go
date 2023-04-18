package entity

//go:generate mockgen -destination=./tests/mocks/repository/account_repository.go -package=mocks AccountRepositoryInterface
type AccountRepositoryInterface interface {
	Save(account *Account) error
	List(limit, offset int) ([]*Account, error)
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}
