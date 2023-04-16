package pkg_validators

type ValidatorInterface interface {
	Validate(dto any) []string
}
