package domain

import (
	pkg_strings "count_on_us/pkg/strings"
	"errors"

	"github.com/google/uuid"
)

type AccountGroup struct {
	ID          string
	Description string
}

func NewAccountGroup(description string) (*AccountGroup, error) {
	accountGroup := &AccountGroup{
		ID:          uuid.NewString(),
		Description: description,
	}
	if err := accountGroup.IsValid(); err != nil {
		return nil, err
	}
	return accountGroup, nil
}

func (a *AccountGroup) IsValid() error {
	if pkg_strings.IsEmpty(a.Description) {
		return errors.New("description is required!")
	}
	return nil
}
