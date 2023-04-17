package repositories

import (
	entity "count_on_us/internal/entity/account"
	"count_on_us/internal/infrastructure/db/sqlx/models"
	"database/sql"
	"fmt"
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
	statement, err := r.DB.Preparex(sql)
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

func (r *AccountRepository) List(limit, offset int) ([]*entity.Account, error) {
	accounts := make([]*entity.Account, 0)
	accountModels := make([]models.Account, 0)

	if limit == 0 {
		limit = 50
	}
	sql := fmt.Sprintf("select * from account limit %d", limit)
	if offset != 0 {
		offset = (offset - 1) * limit
		sql = fmt.Sprintf(`select * from account limit %d offset %d`, limit, offset)
	}
	stmt, err := r.DB.Preparex(sql)
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
	statement, err := r.DB.Preparex("select * from account where id = ?")
	if err != nil {
		return nil, err
	}
	accountDB := models.Account{}
	if err = statement.Get(&accountDB, id); err != nil {
		return nil, err
	}
	return &entity.Account{
		ID:             accountDB.ID,
		DueDate:        accountDB.DueDate,
		PaymentDate:    accountDB.DueDate,
		Description:    accountDB.Description,
		Value:          accountDB.Value,
		Type:           entity.AccountType(accountDB.Type),
		Status:         entity.AccountStatus(accountDB.Status),
		AccountGroupID: accountDB.GroupID,
		OwnerID:        accountDB.OwnerID,
		Installments:   accountDB.Installments,
	}, err
}

func (r *AccountRepository) Update(account *entity.Account) error {
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().String())
	accountModel := models.Account{
		ID:           account.ID,
		Description:  account.Description,
		DueDate:      account.DueDate,
		PaymentDate:  account.PaymentDate,
		Value:        account.Value,
		Type:         string(account.Type),
		Status:       string(account.Status),
		OwnerID:      account.OwnerID,
		GroupID:      account.AccountGroupID,
		Installments: account.Installments,
		UpdatedAt:    &sql.NullTime{Time: updatedAt},
	}

	sql := `
		update account set 
			description=:description,
			due_date=:due_date,
			payment_date=:payment_date,
			value=:value,
			type=:type,
			status=:status,
			owner_id=:owner_id,
			account_group_id=:account_group_id,
			installments=:installments,
			updated_at=:updated_at
		where
			id=:id
	`
	_, err := r.DB.NamedExec(sql, &accountModel)
	if err != nil {
		return err
	}
	return nil
}
