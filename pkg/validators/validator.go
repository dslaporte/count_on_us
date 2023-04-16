package pkg_validators

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() ValidatorInterface {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := validator.RegisterValidation("enum", ValidateEnum); err != nil {
		panic(errors.New("cannot register enum custom validation"))
	}
	if err := validator.RegisterValidation("rfc3339", RFC3339Time); err != nil {
		panic(errors.New("cannot register rfc3339 custom validation"))
	}
	return &Validator{
		validator: validator,
	}
}

func (v *Validator) Validate(dto any) []string {
	errors := make([]string, 0)
	if err := v.validator.Struct(dto); err != nil {
		for _, validationError := range err.(validator.ValidationErrors) {
			errors = append(errors, parseError(validationError))
		}
	}
	return errors

}

func parseError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("Field '%s' is required!", e.Field())
	case "gt":
		return fmt.Sprintf("Field '%s' must be greater than %v", e.Field(), e.Value())
	case "rfc3339":
		return fmt.Sprintf("Field '%s' must be in RFC3339 format (2019-10-12T07:20:50.52Z)", e.Field())
	case "enum":
		return fmt.Sprintf("Invalid value for field '%s'!", e.Field())
	default:
		return ""
	}
}
