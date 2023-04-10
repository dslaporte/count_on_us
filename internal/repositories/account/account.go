package repository

import (
	models "count_on_us/infra/db/models"
	entity "count_on_us/internal/entity/account"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	DB *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) Save(account *entity.Account) error {
	sql := `
		insert into account (
			id,
			description,
			due_date,
			payment_date,
			value,
			type,
			status,
			owner_id,
			account_group_id,
			installments,
			created_at,
			updated_at
		) values (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`
	statement, err := r.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		account.ID,
		account.Description,
		account.DueDate,
		account.PaymentDate,
		account.Value,
		account.Type,
		account.Status,
		account.OwnerID,
		account.AccountGroupID,
		account.Installments,
		time.Now(),
		nil,
	)
	return err
}

func (r *AccountRepository) List() ([]*entity.Account, error) {
	accounts := make([]*entity.Account, 0)
	accountModels := make([]models.Account, 0)
	query := `select * from account`
	stmt, err := r.DB.Preparex(query)
	if err != nil {
		return accounts, err
	}
	if err = stmt.Select(&accountModels); err != nil {
		return accounts, err
	}
	for _, a := range accountModels {
		accounts = append(accounts, &entity.Account{
			ID:             a.ID,
			DueDate:        a.DueDate,
			PaymentDate:    a.PaymentDate,
			Description:    a.Description,
			Value:          a.Value,
			Type:           entity.AccountType(a.Type),
			Status:         entity.AccountStatus(a.Status),
			AccountGroupID: a.GroupID,
			OwnerID:        a.OwnerID,
			Installments:   a.Installments,
		})
	}
	return accounts, err
}

func (r *AccountRepository) FindByID(id string) (*entity.Account, error) {
	return nil, nil
}

func (r *AccountRepository) Update(account *entity.Account) error {
	return nil
}
