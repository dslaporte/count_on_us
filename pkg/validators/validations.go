package pkg_validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EnumValid interface {
	Valid() bool
}

func RFC3339Time(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	_, err := time.Parse(time.RFC3339, timeStr)
	return err == nil
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.Valid()
	}
	return false
}
